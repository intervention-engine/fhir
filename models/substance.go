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

type Substance struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Category       []CodeableConcept              `bson:"category,omitempty" json:"category,omitempty"`
	Code           *CodeableConcept               `bson:"code,omitempty" json:"code,omitempty"`
	Description    string                         `bson:"description,omitempty" json:"description,omitempty"`
	Instance       []SubstanceInstanceComponent   `bson:"instance,omitempty" json:"instance,omitempty"`
	Ingredient     []SubstanceIngredientComponent `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Substance) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Substance"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Substance), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Substance) GetBSON() (interface{}, error) {
	x.ResourceType = "Substance"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "substance" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type substance Substance

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Substance) UnmarshalJSON(data []byte) (err error) {
	x2 := substance{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Substance(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Substance) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Substance"
	} else if x.ResourceType != "Substance" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Substance, instead received %s", x.ResourceType))
	}
	return nil
}

type SubstanceInstanceComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      *Identifier   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Expiry          *FHIRDateTime `bson:"expiry,omitempty" json:"expiry,omitempty"`
	Quantity        *Quantity     `bson:"quantity,omitempty" json:"quantity,omitempty"`
}

type SubstanceIngredientComponent struct {
	BackboneElement `bson:",inline"`
	Quantity        *Ratio     `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Substance       *Reference `bson:"substance,omitempty" json:"substance,omitempty"`
}

type SubstancePlus struct {
	Substance                     `bson:",inline"`
	SubstancePlusRelatedResources `bson:",inline"`
}

type SubstancePlusRelatedResources struct {
	IncludedSubstanceResourcesReferencedBySubstance             *[]Substance             `bson:"_includedSubstanceResourcesReferencedBySubstance,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedSpecimenResourcesReferencingSubject              *[]Specimen              `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationResourcesReferencingIngredient         *[]Medication            `bson:"_revIncludedMedicationResourcesReferencingIngredient,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingSubject                 *[]Order                 `bson:"_revIncludedOrderResourcesReferencingSubject,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedContractResourcesReferencingActor                *[]Contract              `bson:"_revIncludedContractResourcesReferencingActor,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedGroupResourcesReferencingMember                  *[]Group                 `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedSubstanceResourcesReferencingSubstance           *[]Substance             `bson:"_revIncludedSubstanceResourcesReferencingSubstance,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (s *SubstancePlusRelatedResources) GetIncludedSubstanceResourceReferencedBySubstance() (substance *Substance, err error) {
	if s.IncludedSubstanceResourcesReferencedBySubstance == nil {
		err = errors.New("Included substances not requested")
	} else if len(*s.IncludedSubstanceResourcesReferencedBySubstance) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*s.IncludedSubstanceResourcesReferencedBySubstance))
	} else if len(*s.IncludedSubstanceResourcesReferencedBySubstance) == 1 {
		substance = &(*s.IncludedSubstanceResourcesReferencedBySubstance)[0]
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingSubject() (specimen []Specimen, err error) {
	if s.RevIncludedSpecimenResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *s.RevIncludedSpecimenResourcesReferencingSubject
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedMedicationResourcesReferencingIngredient() (medications []Medication, err error) {
	if s.RevIncludedMedicationResourcesReferencingIngredient == nil {
		err = errors.New("RevIncluded medications not requested")
	} else {
		medications = *s.RevIncludedMedicationResourcesReferencingIngredient
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if s.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *s.RevIncludedListResourcesReferencingItem
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedOrderResourcesReferencingSubject() (orders []Order, err error) {
	if s.RevIncludedOrderResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *s.RevIncludedOrderResourcesReferencingSubject
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if s.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *s.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedContractResourcesReferencingActor() (contracts []Contract, err error) {
	if s.RevIncludedContractResourcesReferencingActor == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *s.RevIncludedContractResourcesReferencingActor
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if s.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *s.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if s.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *s.RevIncludedGroupResourcesReferencingMember
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedSubstanceResourcesReferencingSubstance() (substances []Substance, err error) {
	if s.RevIncludedSubstanceResourcesReferencingSubstance == nil {
		err = errors.New("RevIncluded substances not requested")
	} else {
		substances = *s.RevIncludedSubstanceResourcesReferencingSubstance
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if s.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *s.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *s.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *s.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if s.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *s.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *s.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if s.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *s.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (s *SubstancePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedSubstanceResourcesReferencedBySubstance != nil {
		for idx := range *s.IncludedSubstanceResourcesReferencedBySubstance {
			rsc := (*s.IncludedSubstanceResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *SubstancePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*s.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMedicationResourcesReferencingIngredient != nil {
		for idx := range *s.RevIncludedMedicationResourcesReferencingIngredient {
			rsc := (*s.RevIncludedMedicationResourcesReferencingIngredient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedOrderResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedOrderResourcesReferencingSubject {
			rsc := (*s.RevIncludedOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *s.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*s.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingActor {
			rsc := (*s.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *s.RevIncludedGroupResourcesReferencingMember {
			rsc := (*s.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSubstanceResourcesReferencingSubstance != nil {
		for idx := range *s.RevIncludedSubstanceResourcesReferencingSubstance {
			rsc := (*s.RevIncludedSubstanceResourcesReferencingSubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *s.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*s.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*s.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *SubstancePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedSubstanceResourcesReferencedBySubstance != nil {
		for idx := range *s.IncludedSubstanceResourcesReferencedBySubstance {
			rsc := (*s.IncludedSubstanceResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*s.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMedicationResourcesReferencingIngredient != nil {
		for idx := range *s.RevIncludedMedicationResourcesReferencingIngredient {
			rsc := (*s.RevIncludedMedicationResourcesReferencingIngredient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedOrderResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedOrderResourcesReferencingSubject {
			rsc := (*s.RevIncludedOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *s.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*s.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingActor {
			rsc := (*s.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *s.RevIncludedGroupResourcesReferencingMember {
			rsc := (*s.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSubstanceResourcesReferencingSubstance != nil {
		for idx := range *s.RevIncludedSubstanceResourcesReferencingSubstance {
			rsc := (*s.RevIncludedSubstanceResourcesReferencingSubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *s.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*s.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*s.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
