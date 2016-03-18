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

type NutritionOrder struct {
	DomainResource         `bson:",inline"`
	Patient                *Reference                             `bson:"patient,omitempty" json:"patient,omitempty"`
	Orderer                *Reference                             `bson:"orderer,omitempty" json:"orderer,omitempty"`
	Identifier             []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Encounter              *Reference                             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateTime               *FHIRDateTime                          `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Status                 string                                 `bson:"status,omitempty" json:"status,omitempty"`
	AllergyIntolerance     []Reference                            `bson:"allergyIntolerance,omitempty" json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier []CodeableConcept                      `bson:"foodPreferenceModifier,omitempty" json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier    []CodeableConcept                      `bson:"excludeFoodModifier,omitempty" json:"excludeFoodModifier,omitempty"`
	OralDiet               *NutritionOrderOralDietComponent       `bson:"oralDiet,omitempty" json:"oralDiet,omitempty"`
	Supplement             []NutritionOrderSupplementComponent    `bson:"supplement,omitempty" json:"supplement,omitempty"`
	EnteralFormula         *NutritionOrderEnteralFormulaComponent `bson:"enteralFormula,omitempty" json:"enteralFormula,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *NutritionOrder) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "NutritionOrder"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to NutritionOrder), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *NutritionOrder) GetBSON() (interface{}, error) {
	x.ResourceType = "NutritionOrder"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "nutritionOrder" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type nutritionOrder NutritionOrder

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *NutritionOrder) UnmarshalJSON(data []byte) (err error) {
	x2 := nutritionOrder{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = NutritionOrder(x2)
		return x.checkResourceType()
	}
	return
}

func (x *NutritionOrder) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "NutritionOrder"
	} else if x.ResourceType != "NutritionOrder" {
		return errors.New(fmt.Sprintf("Expected resourceType to be NutritionOrder, instead received %s", x.ResourceType))
	}
	return nil
}

type NutritionOrderOralDietComponent struct {
	BackboneElement      `bson:",inline"`
	Type                 []CodeableConcept                         `bson:"type,omitempty" json:"type,omitempty"`
	Schedule             []Timing                                  `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Nutrient             []NutritionOrderOralDietNutrientComponent `bson:"nutrient,omitempty" json:"nutrient,omitempty"`
	Texture              []NutritionOrderOralDietTextureComponent  `bson:"texture,omitempty" json:"texture,omitempty"`
	FluidConsistencyType []CodeableConcept                         `bson:"fluidConsistencyType,omitempty" json:"fluidConsistencyType,omitempty"`
	Instruction          string                                    `bson:"instruction,omitempty" json:"instruction,omitempty"`
}

type NutritionOrderOralDietNutrientComponent struct {
	BackboneElement `bson:",inline"`
	Modifier        *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Amount          *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
}

type NutritionOrderOralDietTextureComponent struct {
	BackboneElement `bson:",inline"`
	Modifier        *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	FoodType        *CodeableConcept `bson:"foodType,omitempty" json:"foodType,omitempty"`
}

type NutritionOrderSupplementComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ProductName     string           `bson:"productName,omitempty" json:"productName,omitempty"`
	Schedule        []Timing         `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity        *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Instruction     string           `bson:"instruction,omitempty" json:"instruction,omitempty"`
}

type NutritionOrderEnteralFormulaComponent struct {
	BackboneElement           `bson:",inline"`
	BaseFormulaType           *CodeableConcept                                      `bson:"baseFormulaType,omitempty" json:"baseFormulaType,omitempty"`
	BaseFormulaProductName    string                                                `bson:"baseFormulaProductName,omitempty" json:"baseFormulaProductName,omitempty"`
	AdditiveType              *CodeableConcept                                      `bson:"additiveType,omitempty" json:"additiveType,omitempty"`
	AdditiveProductName       string                                                `bson:"additiveProductName,omitempty" json:"additiveProductName,omitempty"`
	CaloricDensity            *Quantity                                             `bson:"caloricDensity,omitempty" json:"caloricDensity,omitempty"`
	RouteofAdministration     *CodeableConcept                                      `bson:"routeofAdministration,omitempty" json:"routeofAdministration,omitempty"`
	Administration            []NutritionOrderEnteralFormulaAdministrationComponent `bson:"administration,omitempty" json:"administration,omitempty"`
	MaxVolumeToDeliver        *Quantity                                             `bson:"maxVolumeToDeliver,omitempty" json:"maxVolumeToDeliver,omitempty"`
	AdministrationInstruction string                                                `bson:"administrationInstruction,omitempty" json:"administrationInstruction,omitempty"`
}

type NutritionOrderEnteralFormulaAdministrationComponent struct {
	BackboneElement    `bson:",inline"`
	Schedule           *Timing   `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity           *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	RateSimpleQuantity *Quantity `bson:"rateSimpleQuantity,omitempty" json:"rateSimpleQuantity,omitempty"`
	RateRatio          *Ratio    `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
}

type NutritionOrderPlus struct {
	NutritionOrder                     `bson:",inline"`
	NutritionOrderPlusRelatedResources `bson:",inline"`
}

type NutritionOrderPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByProvider           *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByProvider,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter             *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference    *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAction     *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPlan       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPlan,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (n *NutritionOrderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByProvider() (practitioner *Practitioner, err error) {
	if n.IncludedPractitionerResourcesReferencedByProvider == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*n.IncludedPractitionerResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*n.IncludedPractitionerResourcesReferencedByProvider))
	} else if len(*n.IncludedPractitionerResourcesReferencedByProvider) == 1 {
		practitioner = &(*n.IncludedPractitionerResourcesReferencedByProvider)[0]
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if n.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*n.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*n.IncludedPatientResourcesReferencedByPatient))
	} else if len(*n.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*n.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if n.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*n.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*n.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*n.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*n.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if n.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *n.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if n.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *n.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *n.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if n.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *n.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if n.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *n.RevIncludedListResourcesReferencingItem
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if n.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *n.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if n.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *n.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if n.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *n.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if n.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *n.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if n.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *n.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if n.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *n.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *n.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if n.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *n.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if n.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *n.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if n.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *n.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAction() (clinicalImpressions []ClinicalImpression, err error) {
	if n.RevIncludedClinicalImpressionResourcesReferencingAction == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *n.RevIncludedClinicalImpressionResourcesReferencingAction
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPlan() (clinicalImpressions []ClinicalImpression, err error) {
	if n.RevIncludedClinicalImpressionResourcesReferencingPlan == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *n.RevIncludedClinicalImpressionResourcesReferencingPlan
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if n.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *n.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (n *NutritionOrderPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *n.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*n.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *n.IncludedPatientResourcesReferencedByPatient {
			rsc := (*n.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *n.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*n.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *n.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*n.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedListResourcesReferencingItem {
			rsc := (*n.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *n.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*n.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*n.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *n.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*n.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*n.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*n.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *n.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*n.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *n.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*n.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*n.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *n.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*n.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *n.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*n.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (n *NutritionOrderPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *n.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*n.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *n.IncludedPatientResourcesReferencedByPatient {
			rsc := (*n.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *n.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*n.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *n.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*n.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedListResourcesReferencingItem {
			rsc := (*n.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *n.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*n.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*n.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *n.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*n.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*n.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*n.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *n.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*n.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *n.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*n.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*n.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *n.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*n.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *n.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*n.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
