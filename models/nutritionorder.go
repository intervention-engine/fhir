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

type NutritionOrder struct {
	Id                     string                        `json:"-" bson:"_id"`
	Subject                Reference                     `bson:"subject,omitempty", json:"subject,omitempty"`
	Orderer                Reference                     `bson:"orderer,omitempty", json:"orderer,omitempty"`
	Identifier             []Identifier                  `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Encounter              Reference                     `bson:"encounter,omitempty", json:"encounter,omitempty"`
	DateTime               FHIRDateTime                  `bson:"dateTime,omitempty", json:"dateTime,omitempty"`
	AllergyIntolerance     []Reference                   `bson:"allergyIntolerance,omitempty", json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier []CodeableConcept             `bson:"foodPreferenceModifier,omitempty", json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier    []CodeableConcept             `bson:"excludeFoodModifier,omitempty", json:"excludeFoodModifier,omitempty"`
	Item                   []NutritionOrderItemComponent `bson:"item,omitempty", json:"item,omitempty"`
	Status                 string                        `bson:"status,omitempty", json:"status,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec oralDiet
type NutritionOrderItemOralDietComponent struct {
	Code                   []CodeableConcept `bson:"code,omitempty", json:"code,omitempty"`
	NutrientModifier       []CodeableConcept `bson:"nutrientModifier,omitempty", json:"nutrientModifier,omitempty"`
	NutrientAmountQuantity Quantity          `bson:"nutrientAmountQuantity,omitempty", json:"nutrientAmountQuantity,omitempty"`
	NutrientAmountRange    Range             `bson:"nutrientAmountRange,omitempty", json:"nutrientAmountRange,omitempty"`
	TextureModifier        []CodeableConcept `bson:"textureModifier,omitempty", json:"textureModifier,omitempty"`
	FoodType               []CodeableConcept `bson:"foodType,omitempty", json:"foodType,omitempty"`
	FluidConsistencyType   []CodeableConcept `bson:"fluidConsistencyType,omitempty", json:"fluidConsistencyType,omitempty"`
	Description            string            `bson:"description,omitempty", json:"description,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec supplement
type NutritionOrderItemSupplementComponent struct {
	Type     []CodeableConcept `bson:"type,omitempty", json:"type,omitempty"`
	Quantity Quantity          `bson:"quantity,omitempty", json:"quantity,omitempty"`
	Name     string            `bson:"name,omitempty", json:"name,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec enteralFormula
type NutritionOrderItemEnteralFormulaComponent struct {
	BaseFormulaType       CodeableConcept   `bson:"baseFormulaType,omitempty", json:"baseFormulaType,omitempty"`
	AdditiveType          []CodeableConcept `bson:"additiveType,omitempty", json:"additiveType,omitempty"`
	CaloricDensity        []Quantity        `bson:"caloricDensity,omitempty", json:"caloricDensity,omitempty"`
	RouteofAdministration []CodeableConcept `bson:"routeofAdministration,omitempty", json:"routeofAdministration,omitempty"`
	Rate                  []Quantity        `bson:"rate,omitempty", json:"rate,omitempty"`
	BaseFormulaName       string            `bson:"baseFormulaName,omitempty", json:"baseFormulaName,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec item
type NutritionOrderItemComponent struct {
	ScheduledTiming Timing                                    `bson:"scheduledTiming,omitempty", json:"scheduledTiming,omitempty"`
	ScheduledPeriod Period                                    `bson:"scheduledPeriod,omitempty", json:"scheduledPeriod,omitempty"`
	IsInEffect      bool                                      `bson:"isInEffect,omitempty", json:"isInEffect,omitempty"`
	OralDiet        NutritionOrderItemOralDietComponent       `bson:"oralDiet,omitempty", json:"oralDiet,omitempty"`
	Supplement      NutritionOrderItemSupplementComponent     `bson:"supplement,omitempty", json:"supplement,omitempty"`
	EnteralFormula  NutritionOrderItemEnteralFormulaComponent `bson:"enteralFormula,omitempty", json:"enteralFormula,omitempty"`
}
