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

type ClaimResponse struct {
	DomainResource          `bson:",inline"`
	Identifier              []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Request                 *Reference                        `bson:"request,omitempty" json:"request,omitempty"`
	Ruleset                 *Coding                           `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset         *Coding                           `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created                 *FHIRDateTime                     `bson:"created,omitempty" json:"created,omitempty"`
	Organization            *Reference                        `bson:"organization,omitempty" json:"organization,omitempty"`
	RequestProvider         *Reference                        `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization     *Reference                        `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Outcome                 string                            `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition             string                            `bson:"disposition,omitempty" json:"disposition,omitempty"`
	PayeeType               *Coding                           `bson:"payeeType,omitempty" json:"payeeType,omitempty"`
	Item                    []ClaimResponseItemsComponent     `bson:"item,omitempty" json:"item,omitempty"`
	AddItem                 []ClaimResponseAddedItemComponent `bson:"addItem,omitempty" json:"addItem,omitempty"`
	Error                   []ClaimResponseErrorsComponent    `bson:"error,omitempty" json:"error,omitempty"`
	TotalCost               *Quantity                         `bson:"totalCost,omitempty" json:"totalCost,omitempty"`
	UnallocDeductable       *Quantity                         `bson:"unallocDeductable,omitempty" json:"unallocDeductable,omitempty"`
	TotalBenefit            *Quantity                         `bson:"totalBenefit,omitempty" json:"totalBenefit,omitempty"`
	PaymentAdjustment       *Quantity                         `bson:"paymentAdjustment,omitempty" json:"paymentAdjustment,omitempty"`
	PaymentAdjustmentReason *Coding                           `bson:"paymentAdjustmentReason,omitempty" json:"paymentAdjustmentReason,omitempty"`
	PaymentDate             *FHIRDateTime                     `bson:"paymentDate,omitempty" json:"paymentDate,omitempty"`
	PaymentAmount           *Quantity                         `bson:"paymentAmount,omitempty" json:"paymentAmount,omitempty"`
	PaymentRef              *Identifier                       `bson:"paymentRef,omitempty" json:"paymentRef,omitempty"`
	Reserved                *Coding                           `bson:"reserved,omitempty" json:"reserved,omitempty"`
	Form                    *Coding                           `bson:"form,omitempty" json:"form,omitempty"`
	Note                    []ClaimResponseNotesComponent     `bson:"note,omitempty" json:"note,omitempty"`
	Coverage                []ClaimResponseCoverageComponent  `bson:"coverage,omitempty" json:"coverage,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ClaimResponse) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		ClaimResponse
	}{
		ResourceType:  "ClaimResponse",
		ClaimResponse: *resource,
	}
	return json.Marshal(x)
}

// The "claimResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type claimResponse ClaimResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ClaimResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := claimResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ClaimResponse(x2)
	}
	return
}

type ClaimResponseItemsComponent struct {
	SequenceLinkId *uint32                                  `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	NoteNumber     []uint32                                 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication   []ClaimResponseItemAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail         []ClaimResponseItemDetailComponent       `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ClaimResponseItemAdjudicationComponent struct {
	Code   *Coding   `bson:"code,omitempty" json:"code,omitempty"`
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value  *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ClaimResponseItemDetailComponent struct {
	SequenceLinkId *uint32                                    `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Adjudication   []ClaimResponseDetailAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail      []ClaimResponseSubDetailComponent          `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ClaimResponseDetailAdjudicationComponent struct {
	Code   *Coding   `bson:"code,omitempty" json:"code,omitempty"`
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value  *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ClaimResponseSubDetailComponent struct {
	SequenceLinkId *uint32                                       `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Adjudication   []ClaimResponseSubdetailAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ClaimResponseSubdetailAdjudicationComponent struct {
	Code   *Coding   `bson:"code,omitempty" json:"code,omitempty"`
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value  *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ClaimResponseAddedItemComponent struct {
	SequenceLinkId   []uint32                                      `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Service          *Coding                                       `bson:"service,omitempty" json:"service,omitempty"`
	Fee              *Quantity                                     `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumberLinkId []uint32                                      `bson:"noteNumberLinkId,omitempty" json:"noteNumberLinkId,omitempty"`
	Adjudication     []ClaimResponseAddedItemAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail           []ClaimResponseAddedItemsDetailComponent      `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ClaimResponseAddedItemAdjudicationComponent struct {
	Code   *Coding   `bson:"code,omitempty" json:"code,omitempty"`
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value  *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ClaimResponseAddedItemsDetailComponent struct {
	Service      *Coding                                             `bson:"service,omitempty" json:"service,omitempty"`
	Fee          *Quantity                                           `bson:"fee,omitempty" json:"fee,omitempty"`
	Adjudication []ClaimResponseAddedItemDetailAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ClaimResponseAddedItemDetailAdjudicationComponent struct {
	Code   *Coding   `bson:"code,omitempty" json:"code,omitempty"`
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value  *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ClaimResponseErrorsComponent struct {
	SequenceLinkId          *uint32 `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	DetailSequenceLinkId    *uint32 `bson:"detailSequenceLinkId,omitempty" json:"detailSequenceLinkId,omitempty"`
	SubdetailSequenceLinkId *uint32 `bson:"subdetailSequenceLinkId,omitempty" json:"subdetailSequenceLinkId,omitempty"`
	Code                    *Coding `bson:"code,omitempty" json:"code,omitempty"`
}

type ClaimResponseNotesComponent struct {
	Number *uint32 `bson:"number,omitempty" json:"number,omitempty"`
	Type   *Coding `bson:"type,omitempty" json:"type,omitempty"`
	Text   string  `bson:"text,omitempty" json:"text,omitempty"`
}

type ClaimResponseCoverageComponent struct {
	Sequence            *uint32    `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Focal               *bool      `bson:"focal,omitempty" json:"focal,omitempty"`
	Coverage            *Reference `bson:"coverage,omitempty" json:"coverage,omitempty"`
	BusinessArrangement string     `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	Relationship        *Coding    `bson:"relationship,omitempty" json:"relationship,omitempty"`
	PreAuthRef          []string   `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
	OriginalRuleset     *Coding    `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
}

type ClaimResponsePlus struct {
	ClaimResponse             `bson:",inline"`
	ClaimResponsePlusIncludes `bson:",inline"`
}

type ClaimResponsePlusIncludes struct {
}

func (c *ClaimResponsePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}
