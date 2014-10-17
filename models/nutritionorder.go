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

import "time"

type NutritionOrder struct {
	Id                     string                        `json:"-" bson:"_id"`
	Subject                Reference                     `bson:"subject"`
	Orderer                Reference                     `bson:"orderer"`
	Identifier             []Identifier                  `bson:"identifier"`
	Encounter              Reference                     `bson:"encounter"`
	DateTime               time.Time                     `bson:"dateTime"`
	AllergyIntolerance     []Reference                   `bson:"allergyIntolerance"`
	FoodPreferenceModifier []CodeableConcept             `bson:"foodPreferenceModifier"`
	ExcludeFoodModifier    []CodeableConcept             `bson:"excludeFoodModifier"`
	Item                   []NutritionOrderItemComponent `bson:"item"`
	Status                 string                        `bson:"status"`
}

// This is an ugly hack to deal with embedded structures in the spec oralDiet
type NutritionOrderItemOralDietComponent struct {
	Code                   []CodeableConcept `bson:"code"`
	NutrientModifier       []CodeableConcept `bson:"nutrientModifier"`
	NutrientAmountQuantity Quantity          `bson:"nutrientAmountQuantity"`
	NutrientAmountRange    Range             `bson:"nutrientAmountRange"`
	TextureModifier        []CodeableConcept `bson:"textureModifier"`
	FoodType               []CodeableConcept `bson:"foodType"`
	FluidConsistencyType   []CodeableConcept `bson:"fluidConsistencyType"`
	Description            string            `bson:"description"`
}

// This is an ugly hack to deal with embedded structures in the spec supplement
type NutritionOrderItemSupplementComponent struct {
	FhirType []CodeableConcept `bson:"fhirType"`
	Quantity Quantity          `bson:"quantity"`
	Name     string            `bson:"name"`
}

// This is an ugly hack to deal with embedded structures in the spec enteralFormula
type NutritionOrderItemEnteralFormulaComponent struct {
	BaseFormulaType       CodeableConcept   `bson:"baseFormulaType"`
	AdditiveType          []CodeableConcept `bson:"additiveType"`
	CaloricDensity        []Quantity        `bson:"caloricDensity"`
	RouteofAdministration []CodeableConcept `bson:"routeofAdministration"`
	Rate                  []Quantity        `bson:"rate"`
	BaseFormulaName       string            `bson:"baseFormulaName"`
}

// This is an ugly hack to deal with embedded structures in the spec item
type NutritionOrderItemComponent struct {
	ScheduledTiming Timing                                    `bson:"scheduledTiming"`
	ScheduledPeriod Period                                    `bson:"scheduledPeriod"`
	IsInEffect      bool                                      `bson:"isInEffect"`
	OralDiet        NutritionOrderItemOralDietComponent       `bson:"oralDiet"`
	Supplement      NutritionOrderItemSupplementComponent     `bson:"supplement"`
	EnteralFormula  NutritionOrderItemEnteralFormulaComponent `bson:"enteralFormula"`
}
