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

type Medication struct {
	Id           string                      `json:"-" bson:"_id"`
	Name         string                      `bson:"name,omitempty" json:"name,omitempty"`
	Code         *CodeableConcept            `bson:"code,omitempty" json:"code,omitempty"`
	IsBrand      *bool                       `bson:"isBrand,omitempty" json:"isBrand,omitempty"`
	Manufacturer *Reference                  `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Kind         string                      `bson:"kind,omitempty" json:"kind,omitempty"`
	Product      *MedicationProductComponent `bson:"product,omitempty" json:"product,omitempty"`
	Package      *MedicationPackageComponent `bson:"package,omitempty" json:"package,omitempty"`
}

type MedicationProductComponent struct {
	Form       *CodeableConcept                       `bson:"form,omitempty" json:"form,omitempty"`
	Ingredient []MedicationProductIngredientComponent `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	Batch      []MedicationProductBatchComponent      `bson:"batch,omitempty" json:"batch,omitempty"`
}

type MedicationProductIngredientComponent struct {
	Item   *Reference `bson:"item,omitempty" json:"item,omitempty"`
	Amount *Ratio     `bson:"amount,omitempty" json:"amount,omitempty"`
}

type MedicationProductBatchComponent struct {
	LotNumber      string        `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	ExpirationDate *FHIRDateTime `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
}

type MedicationPackageComponent struct {
	Container *CodeableConcept                    `bson:"container,omitempty" json:"container,omitempty"`
	Content   []MedicationPackageContentComponent `bson:"content,omitempty" json:"content,omitempty"`
}

type MedicationPackageContentComponent struct {
	Item   *Reference `bson:"item,omitempty" json:"item,omitempty"`
	Amount *Quantity  `bson:"amount,omitempty" json:"amount,omitempty"`
}

type MedicationBundle struct {
	Id    string                  `json:"id,omitempty"`
	Type  string                  `json:"resourceType,omitempty"`
	Base  string                  `json:"base,omitempty"`
	Total int                     `json:"total,omitempty"`
	Link  []BundleLinkComponent   `json:"link,omitempty"`
	Entry []MedicationBundleEntry `json:"entry,omitempty"`
}

type MedicationBundleEntry struct {
	Id       string                `json:"id,omitempty"`
	Base     string                `json:"base,omitempty"`
	Link     []BundleLinkComponent `json:"link,omitempty"`
	Resource Medication            `json:"resource,omitempty"`
}

func (resource *Medication) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Medication
	}{
		ResourceType: "Medication",
		Medication:   *resource,
	}
	return json.Marshal(x)
}
