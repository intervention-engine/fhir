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

type Claim struct {
	DomainResource                 `bson:",inline"`
	Type                           string                        `bson:"type,omitempty" json:"type,omitempty"`
	SubType                        []Coding                      `bson:"subType,omitempty" json:"subType,omitempty"`
	Identifier                     []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ruleset                        *Coding                       `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset                *Coding                       `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created                        *FHIRDateTime                 `bson:"created,omitempty" json:"created,omitempty"`
	BillablePeriod                 *Period                       `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	TargetIdentifier               *Identifier                   `bson:"targetIdentifier,omitempty" json:"targetIdentifier,omitempty"`
	TargetReference                *Reference                    `bson:"targetReference,omitempty" json:"targetReference,omitempty"`
	ProviderIdentifier             *Identifier                   `bson:"providerIdentifier,omitempty" json:"providerIdentifier,omitempty"`
	ProviderReference              *Reference                    `bson:"providerReference,omitempty" json:"providerReference,omitempty"`
	OrganizationIdentifier         *Identifier                   `bson:"organizationIdentifier,omitempty" json:"organizationIdentifier,omitempty"`
	OrganizationReference          *Reference                    `bson:"organizationReference,omitempty" json:"organizationReference,omitempty"`
	Use                            string                        `bson:"use,omitempty" json:"use,omitempty"`
	Priority                       *Coding                       `bson:"priority,omitempty" json:"priority,omitempty"`
	FundsReserve                   *Coding                       `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	EntererIdentifier              *Identifier                   `bson:"entererIdentifier,omitempty" json:"entererIdentifier,omitempty"`
	EntererReference               *Reference                    `bson:"entererReference,omitempty" json:"entererReference,omitempty"`
	FacilityIdentifier             *Identifier                   `bson:"facilityIdentifier,omitempty" json:"facilityIdentifier,omitempty"`
	FacilityReference              *Reference                    `bson:"facilityReference,omitempty" json:"facilityReference,omitempty"`
	Related                        []ClaimRelatedClaimsComponent `bson:"related,omitempty" json:"related,omitempty"`
	PrescriptionIdentifier         *Identifier                   `bson:"prescriptionIdentifier,omitempty" json:"prescriptionIdentifier,omitempty"`
	PrescriptionReference          *Reference                    `bson:"prescriptionReference,omitempty" json:"prescriptionReference,omitempty"`
	OriginalPrescriptionIdentifier *Identifier                   `bson:"originalPrescriptionIdentifier,omitempty" json:"originalPrescriptionIdentifier,omitempty"`
	OriginalPrescriptionReference  *Reference                    `bson:"originalPrescriptionReference,omitempty" json:"originalPrescriptionReference,omitempty"`
	Payee                          *ClaimPayeeComponent          `bson:"payee,omitempty" json:"payee,omitempty"`
	ReferralIdentifier             *Identifier                   `bson:"referralIdentifier,omitempty" json:"referralIdentifier,omitempty"`
	ReferralReference              *Reference                    `bson:"referralReference,omitempty" json:"referralReference,omitempty"`
	OccurrenceCode                 []Coding                      `bson:"occurrenceCode,omitempty" json:"occurrenceCode,omitempty"`
	OccurenceSpanCode              []Coding                      `bson:"occurenceSpanCode,omitempty" json:"occurenceSpanCode,omitempty"`
	ValueCode                      []Coding                      `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	Diagnosis                      []ClaimDiagnosisComponent     `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure                      []ClaimProcedureComponent     `bson:"procedure,omitempty" json:"procedure,omitempty"`
	SpecialCondition               []Coding                      `bson:"specialCondition,omitempty" json:"specialCondition,omitempty"`
	PatientIdentifier              *Identifier                   `bson:"patientIdentifier,omitempty" json:"patientIdentifier,omitempty"`
	PatientReference               *Reference                    `bson:"patientReference,omitempty" json:"patientReference,omitempty"`
	Coverage                       []ClaimCoverageComponent      `bson:"coverage,omitempty" json:"coverage,omitempty"`
	AccidentDate                   *FHIRDateTime                 `bson:"accidentDate,omitempty" json:"accidentDate,omitempty"`
	AccidentType                   *Coding                       `bson:"accidentType,omitempty" json:"accidentType,omitempty"`
	AccidentLocationAddress        *Address                      `bson:"accidentLocationAddress,omitempty" json:"accidentLocationAddress,omitempty"`
	AccidentLocationReference      *Reference                    `bson:"accidentLocationReference,omitempty" json:"accidentLocationReference,omitempty"`
	InterventionException          []Coding                      `bson:"interventionException,omitempty" json:"interventionException,omitempty"`
	Onset                          []ClaimOnsetComponent         `bson:"onset,omitempty" json:"onset,omitempty"`
	EmploymentImpacted             *Period                       `bson:"employmentImpacted,omitempty" json:"employmentImpacted,omitempty"`
	Hospitalization                *Period                       `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Item                           []ClaimItemsComponent         `bson:"item,omitempty" json:"item,omitempty"`
	Total                          *Quantity                     `bson:"total,omitempty" json:"total,omitempty"`
	AdditionalMaterial             []Coding                      `bson:"additionalMaterial,omitempty" json:"additionalMaterial,omitempty"`
	MissingTeeth                   []ClaimMissingTeethComponent  `bson:"missingTeeth,omitempty" json:"missingTeeth,omitempty"`
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
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
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

type ClaimRelatedClaimsComponent struct {
	BackboneElement `bson:",inline"`
	ClaimIdentifier *Identifier `bson:"claimIdentifier,omitempty" json:"claimIdentifier,omitempty"`
	ClaimReference  *Reference  `bson:"claimReference,omitempty" json:"claimReference,omitempty"`
	Relationship    *Coding     `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference       *Identifier `bson:"reference,omitempty" json:"reference,omitempty"`
}

type ClaimPayeeComponent struct {
	BackboneElement `bson:",inline"`
	Type            *Coding     `bson:"type,omitempty" json:"type,omitempty"`
	PartyIdentifier *Identifier `bson:"partyIdentifier,omitempty" json:"partyIdentifier,omitempty"`
	PartyReference  *Reference  `bson:"partyReference,omitempty" json:"partyReference,omitempty"`
}

type ClaimDiagnosisComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32 `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Diagnosis       *Coding `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
}

type ClaimProcedureComponent struct {
	BackboneElement    `bson:",inline"`
	Sequence           *uint32       `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Date               *FHIRDateTime `bson:"date,omitempty" json:"date,omitempty"`
	ProcedureCoding    *Coding       `bson:"procedureCoding,omitempty" json:"procedureCoding,omitempty"`
	ProcedureReference *Reference    `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
}

type ClaimCoverageComponent struct {
	BackboneElement     `bson:",inline"`
	Sequence            *uint32     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Focal               *bool       `bson:"focal,omitempty" json:"focal,omitempty"`
	CoverageIdentifier  *Identifier `bson:"coverageIdentifier,omitempty" json:"coverageIdentifier,omitempty"`
	CoverageReference   *Reference  `bson:"coverageReference,omitempty" json:"coverageReference,omitempty"`
	BusinessArrangement string      `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	PreAuthRef          []string    `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference  `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
	OriginalRuleset     *Coding     `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
}

type ClaimOnsetComponent struct {
	BackboneElement `bson:",inline"`
	TimeDate        *FHIRDateTime `bson:"timeDate,omitempty" json:"timeDate,omitempty"`
	TimePeriod      *Period       `bson:"timePeriod,omitempty" json:"timePeriod,omitempty"`
	Type            *Coding       `bson:"type,omitempty" json:"type,omitempty"`
}

type ClaimItemsComponent struct {
	BackboneElement       `bson:",inline"`
	Sequence              *uint32                   `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type                  *Coding                   `bson:"type,omitempty" json:"type,omitempty"`
	ProviderIdentifier    *Identifier               `bson:"providerIdentifier,omitempty" json:"providerIdentifier,omitempty"`
	ProviderReference     *Reference                `bson:"providerReference,omitempty" json:"providerReference,omitempty"`
	SupervisorIdentifier  *Identifier               `bson:"supervisorIdentifier,omitempty" json:"supervisorIdentifier,omitempty"`
	SupervisorReference   *Reference                `bson:"supervisorReference,omitempty" json:"supervisorReference,omitempty"`
	ProviderQualification *Coding                   `bson:"providerQualification,omitempty" json:"providerQualification,omitempty"`
	DiagnosisLinkId       []uint32                  `bson:"diagnosisLinkId,omitempty" json:"diagnosisLinkId,omitempty"`
	Service               *Coding                   `bson:"service,omitempty" json:"service,omitempty"`
	ServiceModifier       []Coding                  `bson:"serviceModifier,omitempty" json:"serviceModifier,omitempty"`
	Modifier              []Coding                  `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode           []Coding                  `bson:"programCode,omitempty" json:"programCode,omitempty"`
	ServicedDate          *FHIRDateTime             `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod        *Period                   `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	Place                 *Coding                   `bson:"place,omitempty" json:"place,omitempty"`
	Quantity              *Quantity                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice             *Quantity                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                *float64                  `bson:"factor,omitempty" json:"factor,omitempty"`
	Points                *float64                  `bson:"points,omitempty" json:"points,omitempty"`
	Net                   *Quantity                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi                   []Reference               `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite              *Coding                   `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite               []Coding                  `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Detail                []ClaimDetailComponent    `bson:"detail,omitempty" json:"detail,omitempty"`
	Prosthesis            *ClaimProsthesisComponent `bson:"prosthesis,omitempty" json:"prosthesis,omitempty"`
}

type ClaimDetailComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32                   `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type            *Coding                   `bson:"type,omitempty" json:"type,omitempty"`
	Service         *Coding                   `bson:"service,omitempty" json:"service,omitempty"`
	ProgramCode     []Coding                  `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity        *Quantity                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64                  `bson:"factor,omitempty" json:"factor,omitempty"`
	Points          *float64                  `bson:"points,omitempty" json:"points,omitempty"`
	Net             *Quantity                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi             []Reference               `bson:"udi,omitempty" json:"udi,omitempty"`
	SubDetail       []ClaimSubDetailComponent `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ClaimSubDetailComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type            *Coding     `bson:"type,omitempty" json:"type,omitempty"`
	Service         *Coding     `bson:"service,omitempty" json:"service,omitempty"`
	ProgramCode     []Coding    `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity        *Quantity   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity   `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64    `bson:"factor,omitempty" json:"factor,omitempty"`
	Points          *float64    `bson:"points,omitempty" json:"points,omitempty"`
	Net             *Quantity   `bson:"net,omitempty" json:"net,omitempty"`
	Udi             []Reference `bson:"udi,omitempty" json:"udi,omitempty"`
}

type ClaimProsthesisComponent struct {
	BackboneElement `bson:",inline"`
	Initial         *bool         `bson:"initial,omitempty" json:"initial,omitempty"`
	PriorDate       *FHIRDateTime `bson:"priorDate,omitempty" json:"priorDate,omitempty"`
	PriorMaterial   *Coding       `bson:"priorMaterial,omitempty" json:"priorMaterial,omitempty"`
}

type ClaimMissingTeethComponent struct {
	BackboneElement `bson:",inline"`
	Tooth           *Coding       `bson:"tooth,omitempty" json:"tooth,omitempty"`
	Reason          *Coding       `bson:"reason,omitempty" json:"reason,omitempty"`
	ExtractionDate  *FHIRDateTime `bson:"extractionDate,omitempty" json:"extractionDate,omitempty"`
}

type ClaimPlus struct {
	Claim                     `bson:",inline"`
	ClaimPlusRelatedResources `bson:",inline"`
}

type ClaimPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatientreference              *[]Patient               `bson:"_includedPatientResourcesReferencedByPatientreference,omitempty"`
	IncludedPractitionerResourcesReferencedByProviderreference        *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByProviderreference,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganizationreference    *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganizationreference,omitempty"`
	IncludedLocationResourcesReferencedByFacilityreference            *[]Location              `bson:"_includedLocationResourcesReferencedByFacilityreference,omitempty"`
	IncludedOrganizationResourcesReferencedByTargetreference          *[]Organization          `bson:"_includedOrganizationResourcesReferencedByTargetreference,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref         *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref         *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref        *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                     *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                    *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                      *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponsereference     *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponsereference,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequestreference      *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequestreference,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource        *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment           *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                  *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                   *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                        *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingClaimreference *[]ExplanationOfBenefit  `bson:"_revIncludedExplanationOfBenefitResourcesReferencingClaimreference,omitempty"`
	RevIncludedListResourcesReferencingItem                           *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                        *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                       *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedClaimResponseResourcesReferencingRequestreference      *[]ClaimResponse         `bson:"_revIncludedClaimResponseResourcesReferencingRequestreference,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                   *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                 *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                   *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated            *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject       *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestreference    *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequestreference,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger          *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
}

func (c *ClaimPlusRelatedResources) GetIncludedPatientResourceReferencedByPatientreference() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatientreference == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPatientreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPatientreference))
	} else if len(*c.IncludedPatientResourcesReferencedByPatientreference) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPatientreference)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerResourceReferencedByProviderreference() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByProviderreference == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByProviderreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByProviderreference))
	} else if len(*c.IncludedPractitionerResourcesReferencedByProviderreference) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByProviderreference)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganizationreference() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByOrganizationreference == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByOrganizationreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByOrganizationreference))
	} else if len(*c.IncludedOrganizationResourcesReferencedByOrganizationreference) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByOrganizationreference)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedLocationResourceReferencedByFacilityreference() (location *Location, err error) {
	if c.IncludedLocationResourcesReferencedByFacilityreference == nil {
		err = errors.New("Included locations not requested")
	} else if len(*c.IncludedLocationResourcesReferencedByFacilityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*c.IncludedLocationResourcesReferencedByFacilityreference))
	} else if len(*c.IncludedLocationResourcesReferencedByFacilityreference) == 1 {
		location = &(*c.IncludedLocationResourcesReferencedByFacilityreference)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedOrganizationResourceReferencedByTargetreference() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByTargetreference == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByTargetreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByTargetreference))
	} else if len(*c.IncludedOrganizationResourcesReferencedByTargetreference) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByTargetreference)[0]
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

func (c *ClaimPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTtopic
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

func (c *ClaimPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference
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

func (c *ClaimPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
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

func (c *ClaimPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingClaimreference() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if c.RevIncludedExplanationOfBenefitResourcesReferencingClaimreference == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *c.RevIncludedExplanationOfBenefitResourcesReferencingClaimreference
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

func (c *ClaimPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
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

func (c *ClaimPlusRelatedResources) GetRevIncludedClaimResponseResourcesReferencingRequestreference() (claimResponses []ClaimResponse, err error) {
	if c.RevIncludedClaimResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded claimResponses not requested")
	} else {
		claimResponses = *c.RevIncludedClaimResponseResourcesReferencingRequestreference
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

func (c *ClaimPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPatientResourcesReferencedByPatientreference != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatientreference {
			rsc := (*c.IncludedPatientResourcesReferencedByPatientreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByProviderreference != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByProviderreference {
			rsc := (*c.IncludedPractitionerResourcesReferencedByProviderreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByOrganizationreference != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByOrganizationreference {
			rsc := (*c.IncludedOrganizationResourcesReferencedByOrganizationreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByFacilityreference != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByFacilityreference {
			rsc := (*c.IncludedLocationResourcesReferencedByFacilityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByTargetreference != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByTargetreference {
			rsc := (*c.IncludedOrganizationResourcesReferencedByTargetreference)[idx]
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
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*c.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
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
	if c.RevIncludedExplanationOfBenefitResourcesReferencingClaimreference != nil {
		for idx := range *c.RevIncludedExplanationOfBenefitResourcesReferencingClaimreference {
			rsc := (*c.RevIncludedExplanationOfBenefitResourcesReferencingClaimreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *c.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*c.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClaimResponseResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedClaimResponseResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedClaimResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
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
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ClaimPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPatientResourcesReferencedByPatientreference != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatientreference {
			rsc := (*c.IncludedPatientResourcesReferencedByPatientreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByProviderreference != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByProviderreference {
			rsc := (*c.IncludedPractitionerResourcesReferencedByProviderreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByOrganizationreference != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByOrganizationreference {
			rsc := (*c.IncludedOrganizationResourcesReferencedByOrganizationreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByFacilityreference != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByFacilityreference {
			rsc := (*c.IncludedLocationResourcesReferencedByFacilityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByTargetreference != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByTargetreference {
			rsc := (*c.IncludedOrganizationResourcesReferencedByTargetreference)[idx]
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
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*c.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
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
	if c.RevIncludedExplanationOfBenefitResourcesReferencingClaimreference != nil {
		for idx := range *c.RevIncludedExplanationOfBenefitResourcesReferencingClaimreference {
			rsc := (*c.RevIncludedExplanationOfBenefitResourcesReferencingClaimreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *c.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*c.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClaimResponseResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedClaimResponseResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedClaimResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
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
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
