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

import "encoding/json"

type Claim struct {
	DomainResource        `bson:",inline"`
	Type                  string                       `bson:"type,omitempty" json:"type,omitempty"`
	Identifier            []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ruleset               *Coding                      `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset       *Coding                      `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created               *FHIRDateTime                `bson:"created,omitempty" json:"created,omitempty"`
	Target                *Reference                   `bson:"target,omitempty" json:"target,omitempty"`
	Provider              *Reference                   `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization          *Reference                   `bson:"organization,omitempty" json:"organization,omitempty"`
	Use                   string                       `bson:"use,omitempty" json:"use,omitempty"`
	Priority              *Coding                      `bson:"priority,omitempty" json:"priority,omitempty"`
	FundsReserve          *Coding                      `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	Enterer               *Reference                   `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Facility              *Reference                   `bson:"facility,omitempty" json:"facility,omitempty"`
	Prescription          *Reference                   `bson:"prescription,omitempty" json:"prescription,omitempty"`
	OriginalPrescription  *Reference                   `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                 *ClaimPayeeComponent         `bson:"payee,omitempty" json:"payee,omitempty"`
	Referral              *Reference                   `bson:"referral,omitempty" json:"referral,omitempty"`
	Diagnosis             []ClaimDiagnosisComponent    `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Condition             []Coding                     `bson:"condition,omitempty" json:"condition,omitempty"`
	Patient               *Reference                   `bson:"patient,omitempty" json:"patient,omitempty"`
	Coverage              []ClaimCoverageComponent     `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Exception             []Coding                     `bson:"exception,omitempty" json:"exception,omitempty"`
	School                string                       `bson:"school,omitempty" json:"school,omitempty"`
	Accident              *FHIRDateTime                `bson:"accident,omitempty" json:"accident,omitempty"`
	AccidentType          *Coding                      `bson:"accidentType,omitempty" json:"accidentType,omitempty"`
	InterventionException []Coding                     `bson:"interventionException,omitempty" json:"interventionException,omitempty"`
	Item                  []ClaimItemsComponent        `bson:"item,omitempty" json:"item,omitempty"`
	AdditionalMaterials   []Coding                     `bson:"additionalMaterials,omitempty" json:"additionalMaterials,omitempty"`
	MissingTeeth          []ClaimMissingTeethComponent `bson:"missingTeeth,omitempty" json:"missingTeeth,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Claim) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Claim
	}{
		ResourceType: "Claim",
		Claim:        *resource,
	}
	return json.Marshal(x)
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
	}
	return
}

type ClaimPayeeComponent struct {
	Type         *Coding    `bson:"type,omitempty" json:"type,omitempty"`
	Provider     *Reference `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization *Reference `bson:"organization,omitempty" json:"organization,omitempty"`
	Person       *Reference `bson:"person,omitempty" json:"person,omitempty"`
}

type ClaimDiagnosisComponent struct {
	Sequence  *uint32 `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Diagnosis *Coding `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
}

type ClaimCoverageComponent struct {
	Sequence            *uint32    `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Focal               *bool      `bson:"focal,omitempty" json:"focal,omitempty"`
	Coverage            *Reference `bson:"coverage,omitempty" json:"coverage,omitempty"`
	BusinessArrangement string     `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	Relationship        *Coding    `bson:"relationship,omitempty" json:"relationship,omitempty"`
	PreAuthRef          []string   `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
	OriginalRuleset     *Coding    `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
}

type ClaimItemsComponent struct {
	Sequence        *uint32                   `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type            *Coding                   `bson:"type,omitempty" json:"type,omitempty"`
	Provider        *Reference                `bson:"provider,omitempty" json:"provider,omitempty"`
	DiagnosisLinkId []uint32                  `bson:"diagnosisLinkId,omitempty" json:"diagnosisLinkId,omitempty"`
	Service         *Coding                   `bson:"service,omitempty" json:"service,omitempty"`
	ServiceDate     *FHIRDateTime             `bson:"serviceDate,omitempty" json:"serviceDate,omitempty"`
	Quantity        *Quantity                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64                  `bson:"factor,omitempty" json:"factor,omitempty"`
	Points          *float64                  `bson:"points,omitempty" json:"points,omitempty"`
	Net             *Quantity                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi             *Coding                   `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite        *Coding                   `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite         []Coding                  `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Modifier        []Coding                  `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Detail          []ClaimDetailComponent    `bson:"detail,omitempty" json:"detail,omitempty"`
	Prosthesis      *ClaimProsthesisComponent `bson:"prosthesis,omitempty" json:"prosthesis,omitempty"`
}

type ClaimDetailComponent struct {
	Sequence  *uint32                   `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type      *Coding                   `bson:"type,omitempty" json:"type,omitempty"`
	Service   *Coding                   `bson:"service,omitempty" json:"service,omitempty"`
	Quantity  *Quantity                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice *Quantity                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor    *float64                  `bson:"factor,omitempty" json:"factor,omitempty"`
	Points    *float64                  `bson:"points,omitempty" json:"points,omitempty"`
	Net       *Quantity                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi       *Coding                   `bson:"udi,omitempty" json:"udi,omitempty"`
	SubDetail []ClaimSubDetailComponent `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ClaimSubDetailComponent struct {
	Sequence  *uint32   `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type      *Coding   `bson:"type,omitempty" json:"type,omitempty"`
	Service   *Coding   `bson:"service,omitempty" json:"service,omitempty"`
	Quantity  *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice *Quantity `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor    *float64  `bson:"factor,omitempty" json:"factor,omitempty"`
	Points    *float64  `bson:"points,omitempty" json:"points,omitempty"`
	Net       *Quantity `bson:"net,omitempty" json:"net,omitempty"`
	Udi       *Coding   `bson:"udi,omitempty" json:"udi,omitempty"`
}

type ClaimProsthesisComponent struct {
	Initial       *bool         `bson:"initial,omitempty" json:"initial,omitempty"`
	PriorDate     *FHIRDateTime `bson:"priorDate,omitempty" json:"priorDate,omitempty"`
	PriorMaterial *Coding       `bson:"priorMaterial,omitempty" json:"priorMaterial,omitempty"`
}

type ClaimMissingTeethComponent struct {
	Tooth          *Coding       `bson:"tooth,omitempty" json:"tooth,omitempty"`
	Reason         *Coding       `bson:"reason,omitempty" json:"reason,omitempty"`
	ExtractionDate *FHIRDateTime `bson:"extractionDate,omitempty" json:"extractionDate,omitempty"`
}
