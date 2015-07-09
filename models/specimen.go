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

type Specimen struct {
	Id                  string                       `json:"-" bson:"_id"`
	Identifier          []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                *CodeableConcept             `bson:"type,omitempty" json:"type,omitempty"`
	Parent              []Reference                  `bson:"parent,omitempty" json:"parent,omitempty"`
	Subject             *Reference                   `bson:"subject,omitempty" json:"subject,omitempty"`
	AccessionIdentifier *Identifier                  `bson:"accessionIdentifier,omitempty" json:"accessionIdentifier,omitempty"`
	ReceivedTime        *FHIRDateTime                `bson:"receivedTime,omitempty" json:"receivedTime,omitempty"`
	Collection          *SpecimenCollectionComponent `bson:"collection,omitempty" json:"collection,omitempty"`
	Treatment           []SpecimenTreatmentComponent `bson:"treatment,omitempty" json:"treatment,omitempty"`
	Container           []SpecimenContainerComponent `bson:"container,omitempty" json:"container,omitempty"`
}

type SpecimenCollectionComponent struct {
	Collector               *Reference       `bson:"collector,omitempty" json:"collector,omitempty"`
	Comment                 []string         `bson:"comment,omitempty" json:"comment,omitempty"`
	CollectedDateTime       *FHIRDateTime    `bson:"collectedDateTime,omitempty" json:"collectedDateTime,omitempty"`
	CollectedPeriod         *Period          `bson:"collectedPeriod,omitempty" json:"collectedPeriod,omitempty"`
	Quantity                *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Method                  *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	BodySiteCodeableConcept *CodeableConcept `bson:"bodySiteCodeableConcept,omitempty" json:"bodySiteCodeableConcept,omitempty"`
	BodySiteReference       *Reference       `bson:"bodySiteReference,omitempty" json:"bodySiteReference,omitempty"`
}

type SpecimenTreatmentComponent struct {
	Description string           `bson:"description,omitempty" json:"description,omitempty"`
	Procedure   *CodeableConcept `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Additive    []Reference      `bson:"additive,omitempty" json:"additive,omitempty"`
}

type SpecimenContainerComponent struct {
	Identifier              []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Description             string           `bson:"description,omitempty" json:"description,omitempty"`
	Type                    *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Capacity                *Quantity        `bson:"capacity,omitempty" json:"capacity,omitempty"`
	SpecimenQuantity        *Quantity        `bson:"specimenQuantity,omitempty" json:"specimenQuantity,omitempty"`
	AdditiveCodeableConcept *CodeableConcept `bson:"additiveCodeableConcept,omitempty" json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *Reference       `bson:"additiveReference,omitempty" json:"additiveReference,omitempty"`
}

type SpecimenBundle struct {
	Id    string                `json:"id,omitempty"`
	Type  string                `json:"resourceType,omitempty"`
	Base  string                `json:"base,omitempty"`
	Total int                   `json:"total,omitempty"`
	Link  []BundleLinkComponent `json:"link,omitempty"`
	Entry []SpecimenBundleEntry `json:"entry,omitempty"`
}

type SpecimenBundleEntry struct {
	Id       string                `json:"id,omitempty"`
	Base     string                `json:"base,omitempty"`
	Link     []BundleLinkComponent `json:"link,omitempty"`
	Resource Specimen              `json:"resource,omitempty"`
}

func (resource *Specimen) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Specimen
	}{
		ResourceType: "Specimen",
		Specimen:     *resource,
	}
	return json.Marshal(x)
}
