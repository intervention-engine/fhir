// Copyright (c) 2011-2014, HL7, Inc & The MITRE Corporation
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

type Medication struct {
	Id           string                     `json:"-" bson:"_id"`
	Name         string                     `bson:"name"`
	Code         CodeableConcept            `bson:"code"`
	IsBrand      bool                       `bson:"isBrand"`
	Manufacturer Reference                  `bson:"manufacturer"`
	Kind         string                     `bson:"kind"`
	Product      MedicationProductComponent `bson:"product"`
	Package      MedicationPackageComponent `bson:"package"`
}

// This is an ugly hack to deal with embedded structures in the spec ingredient
type MedicationProductIngredientComponent struct {
	Item   Reference `bson:"item"`
	Amount Ratio     `bson:"amount"`
}

// This is an ugly hack to deal with embedded structures in the spec product
type MedicationProductComponent struct {
	Form       CodeableConcept                        `bson:"form"`
	Ingredient []MedicationProductIngredientComponent `bson:"ingredient"`
}

// This is an ugly hack to deal with embedded structures in the spec content
type MedicationPackageContentComponent struct {
	Item   Reference `bson:"item"`
	Amount Quantity  `bson:"amount"`
}

// This is an ugly hack to deal with embedded structures in the spec package
type MedicationPackageComponent struct {
	Container CodeableConcept                     `bson:"container"`
	Content   []MedicationPackageContentComponent `bson:"content"`
}
