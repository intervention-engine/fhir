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

type Medication struct {
	DomainResource `bson:",inline"`
	Code           *CodeableConcept            `bson:"code,omitempty" json:"code,omitempty"`
	IsBrand        *bool                       `bson:"isBrand,omitempty" json:"isBrand,omitempty"`
	Manufacturer   *Reference                  `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Product        *MedicationProductComponent `bson:"product,omitempty" json:"product,omitempty"`
	Package        *MedicationPackageComponent `bson:"package,omitempty" json:"package,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
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

// The "medication" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medication Medication

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Medication) UnmarshalJSON(data []byte) (err error) {
	x2 := medication{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Medication(x2)
	}
	return
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

type MedicationPlus struct {
	Medication             `bson:",inline"`
	MedicationPlusIncludes `bson:",inline"`
}

type MedicationPlusIncludes struct {
	IncludedIngredientMedicationResources *[]Medication   `bson:"_includedIngredientMedicationResources,omitempty"`
	IncludedIngredientSubstanceResources  *[]Substance    `bson:"_includedIngredientSubstanceResources,omitempty"`
	IncludedContentResources              *[]Medication   `bson:"_includedContentResources,omitempty"`
	IncludedManufacturerResources         *[]Organization `bson:"_includedManufacturerResources,omitempty"`
}

func (m *MedicationPlusIncludes) GetIncludedIngredientMedicationResource() (medication *Medication, err error) {
	if m.IncludedIngredientMedicationResources == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedIngredientMedicationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedIngredientMedicationResources))
	} else if len(*m.IncludedIngredientMedicationResources) == 1 {
		medication = &(*m.IncludedIngredientMedicationResources)[0]
	}
	return
}

func (m *MedicationPlusIncludes) GetIncludedIngredientSubstanceResource() (substance *Substance, err error) {
	if m.IncludedIngredientSubstanceResources == nil {
		err = errors.New("Included substances not requested")
	} else if len(*m.IncludedIngredientSubstanceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*m.IncludedIngredientSubstanceResources))
	} else if len(*m.IncludedIngredientSubstanceResources) == 1 {
		substance = &(*m.IncludedIngredientSubstanceResources)[0]
	}
	return
}

func (m *MedicationPlusIncludes) GetIncludedContentResource() (medication *Medication, err error) {
	if m.IncludedContentResources == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedContentResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedContentResources))
	} else if len(*m.IncludedContentResources) == 1 {
		medication = &(*m.IncludedContentResources)[0]
	}
	return
}

func (m *MedicationPlusIncludes) GetIncludedManufacturerResource() (organization *Organization, err error) {
	if m.IncludedManufacturerResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*m.IncludedManufacturerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*m.IncludedManufacturerResources))
	} else if len(*m.IncludedManufacturerResources) == 1 {
		organization = &(*m.IncludedManufacturerResources)[0]
	}
	return
}

func (m *MedicationPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedIngredientMedicationResources != nil {
		for _, r := range *m.IncludedIngredientMedicationResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedIngredientSubstanceResources != nil {
		for _, r := range *m.IncludedIngredientSubstanceResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedContentResources != nil {
		for _, r := range *m.IncludedContentResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedManufacturerResources != nil {
		for _, r := range *m.IncludedManufacturerResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
