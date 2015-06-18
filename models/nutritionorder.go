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

import "time"

type NutritionOrder struct {
	Id                     string                                               `json:"-" bson:"_id"`
	Patient                *Reference                                           `bson:"patient,omitempty" json:"patient,omitempty"`
	Orderer                *Reference                                           `bson:"orderer,omitempty" json:"orderer,omitempty"`
	Identifier             []Identifier                                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Encounter              *Reference                                           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateTime               *FHIRDateTime                                        `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Status                 string                                               `bson:"status,omitempty" json:"status,omitempty"`
	AllergyIntolerance     []Reference                                          `bson:"allergyIntolerance,omitempty" json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier []CodeableConcept                                    `bson:"foodPreferenceModifier,omitempty" json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier    []CodeableConcept                                    `bson:"excludeFoodModifier,omitempty" json:"excludeFoodModifier,omitempty"`
	OralDiet               *NutritionOrderNutritionOrderOralDietComponent       `bson:"oralDiet,omitempty" json:"oralDiet,omitempty"`
	Supplement             []NutritionOrderNutritionOrderSupplementComponent    `bson:"supplement,omitempty" json:"supplement,omitempty"`
	EnteralFormula         *NutritionOrderNutritionOrderEnteralFormulaComponent `bson:"enteralFormula,omitempty" json:"enteralFormula,omitempty"`
}
type NutritionOrderNutritionOrderOralDietComponent struct {
	Type                 []CodeableConcept                                       `bson:"type,omitempty" json:"type,omitempty"`
	Schedule             []Timing                                                `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Nutrient             []NutritionOrderNutritionOrderOralDietNutrientComponent `bson:"nutrient,omitempty" json:"nutrient,omitempty"`
	Texture              []NutritionOrderNutritionOrderOralDietTextureComponent  `bson:"texture,omitempty" json:"texture,omitempty"`
	FluidConsistencyType []CodeableConcept                                       `bson:"fluidConsistencyType,omitempty" json:"fluidConsistencyType,omitempty"`
	Instruction          string                                                  `bson:"instruction,omitempty" json:"instruction,omitempty"`
}
type NutritionOrderNutritionOrderOralDietNutrientComponent struct {
	Modifier *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Amount   *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
}
type NutritionOrderNutritionOrderOralDietTextureComponent struct {
	Modifier *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	FoodType *CodeableConcept `bson:"foodType,omitempty" json:"foodType,omitempty"`
}
type NutritionOrderNutritionOrderSupplementComponent struct {
	Type        *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ProductName string           `bson:"productName,omitempty" json:"productName,omitempty"`
	Schedule    []Timing         `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity    *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Instruction string           `bson:"instruction,omitempty" json:"instruction,omitempty"`
}
type NutritionOrderNutritionOrderEnteralFormulaComponent struct {
	BaseFormulaType           *CodeableConcept                                                    `bson:"baseFormulaType,omitempty" json:"baseFormulaType,omitempty"`
	BaseFormulaProductName    string                                                              `bson:"baseFormulaProductName,omitempty" json:"baseFormulaProductName,omitempty"`
	AdditiveType              *CodeableConcept                                                    `bson:"additiveType,omitempty" json:"additiveType,omitempty"`
	AdditiveProductName       string                                                              `bson:"additiveProductName,omitempty" json:"additiveProductName,omitempty"`
	CaloricDensity            *Quantity                                                           `bson:"caloricDensity,omitempty" json:"caloricDensity,omitempty"`
	RouteofAdministration     *CodeableConcept                                                    `bson:"routeofAdministration,omitempty" json:"routeofAdministration,omitempty"`
	Administration            []NutritionOrderNutritionOrderEnteralFormulaAdministrationComponent `bson:"administration,omitempty" json:"administration,omitempty"`
	MaxVolumeToDeliver        *Quantity                                                           `bson:"maxVolumeToDeliver,omitempty" json:"maxVolumeToDeliver,omitempty"`
	AdministrationInstruction string                                                              `bson:"administrationInstruction,omitempty" json:"administrationInstruction,omitempty"`
}
type NutritionOrderNutritionOrderEnteralFormulaAdministrationComponent struct {
	Schedule     *Timing   `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity     *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	RateQuantity *Quantity `bson:"rateQuantity,omitempty" json:"rateQuantity,omitempty"`
	RateRatio    *Ratio    `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
}

type NutritionOrderBundle struct {
	Type         string                      `json:"resourceType,omitempty"`
	Title        string                      `json:"title,omitempty"`
	Id           string                      `json:"id,omitempty"`
	Updated      time.Time                   `json:"updated,omitempty"`
	TotalResults int                         `json:"totalResults,omitempty"`
	Entry        []NutritionOrderBundleEntry `json:"entry,omitempty"`
	Category     NutritionOrderCategory      `json:"category,omitempty"`
}

type NutritionOrderBundleEntry struct {
	Title    string                 `json:"title,omitempty"`
	Id       string                 `json:"id,omitempty"`
	Content  NutritionOrder         `json:"content,omitempty"`
	Category NutritionOrderCategory `json:"category,omitempty"`
}

type NutritionOrderCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
