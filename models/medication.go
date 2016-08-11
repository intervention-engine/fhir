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
	resource.ResourceType = "Medication"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Medication), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Medication) GetBSON() (interface{}, error) {
	x.ResourceType = "Medication"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
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
		return x.checkResourceType()
	}
	return
}

func (x *Medication) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Medication"
	} else if x.ResourceType != "Medication" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Medication, instead received %s", x.ResourceType))
	}
	return nil
}

type MedicationProductComponent struct {
	BackboneElement `bson:",inline"`
	Form            *CodeableConcept                       `bson:"form,omitempty" json:"form,omitempty"`
	Ingredient      []MedicationProductIngredientComponent `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	Batch           []MedicationProductBatchComponent      `bson:"batch,omitempty" json:"batch,omitempty"`
}

type MedicationProductIngredientComponent struct {
	BackboneElement     `bson:",inline"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
	Amount              *Ratio           `bson:"amount,omitempty" json:"amount,omitempty"`
}

type MedicationProductBatchComponent struct {
	BackboneElement `bson:",inline"`
	LotNumber       string        `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	ExpirationDate  *FHIRDateTime `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
}

type MedicationPackageComponent struct {
	BackboneElement `bson:",inline"`
	Container       *CodeableConcept                    `bson:"container,omitempty" json:"container,omitempty"`
	Content         []MedicationPackageContentComponent `bson:"content,omitempty" json:"content,omitempty"`
}

type MedicationPackageContentComponent struct {
	BackboneElement     `bson:",inline"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
	Amount              *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
}

type MedicationPlus struct {
	Medication                     `bson:",inline"`
	MedicationPlusRelatedResources `bson:",inline"`
}

type MedicationPlusRelatedResources struct {
	IncludedMedicationResourcesReferencedByPackageitem                *[]Medication               `bson:"_includedMedicationResourcesReferencedByPackageitem,omitempty"`
	IncludedMedicationResourcesReferencedByIngredient                 *[]Medication               `bson:"_includedMedicationResourcesReferencedByIngredient,omitempty"`
	IncludedSubstanceResourcesReferencedByIngredient                  *[]Substance                `bson:"_includedSubstanceResourcesReferencedByIngredient,omitempty"`
	IncludedOrganizationResourcesReferencedByManufacturer             *[]Organization             `bson:"_includedOrganizationResourcesReferencedByManufacturer,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref         *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref         *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedMedicationResourcesReferencingPackageitem              *[]Medication               `bson:"_revIncludedMedicationResourcesReferencingPackageitem,omitempty"`
	RevIncludedMedicationResourcesReferencingIngredient               *[]Medication               `bson:"_revIncludedMedicationResourcesReferencingIngredient,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref        *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                     *[]Contract                 `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                    *[]Contract                 `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                      *[]Contract                 `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedGroupResourcesReferencingMember                        *[]Group                    `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponsereference     *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingResponsereference,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequestreference      *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingRequestreference,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource        *[]ImplementationGuide      `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedProtocolResourcesReferencingSubject                    *[]Protocol                 `bson:"_revIncludedProtocolResourcesReferencingSubject,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment           *[]OrderResponse            `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                  *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                   *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                        *[]Task                     `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingItem                           *[]List                     `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                        *[]Order                    `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingMedication *[]MedicationAdministration `bson:"_revIncludedMedicationAdministrationResourcesReferencingMedication,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingMedication      *[]MedicationStatement      `bson:"_revIncludedMedicationStatementResourcesReferencingMedication,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                       *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingMedication       *[]MedicationDispense       `bson:"_revIncludedMedicationDispenseResourcesReferencingMedication,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                   *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedMedicationOrderResourcesReferencingMedication          *[]MedicationOrder          `bson:"_revIncludedMedicationOrderResourcesReferencingMedication,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                 *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                   *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated            *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject       *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestreference    *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequestreference,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger          *[]ClinicalImpression       `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
}

func (m *MedicationPlusRelatedResources) GetIncludedMedicationResourceReferencedByPackageitem() (medication *Medication, err error) {
	if m.IncludedMedicationResourcesReferencedByPackageitem == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedMedicationResourcesReferencedByPackageitem) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedMedicationResourcesReferencedByPackageitem))
	} else if len(*m.IncludedMedicationResourcesReferencedByPackageitem) == 1 {
		medication = &(*m.IncludedMedicationResourcesReferencedByPackageitem)[0]
	}
	return
}

func (m *MedicationPlusRelatedResources) GetIncludedMedicationResourceReferencedByIngredient() (medication *Medication, err error) {
	if m.IncludedMedicationResourcesReferencedByIngredient == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedMedicationResourcesReferencedByIngredient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedMedicationResourcesReferencedByIngredient))
	} else if len(*m.IncludedMedicationResourcesReferencedByIngredient) == 1 {
		medication = &(*m.IncludedMedicationResourcesReferencedByIngredient)[0]
	}
	return
}

func (m *MedicationPlusRelatedResources) GetIncludedSubstanceResourceReferencedByIngredient() (substance *Substance, err error) {
	if m.IncludedSubstanceResourcesReferencedByIngredient == nil {
		err = errors.New("Included substances not requested")
	} else if len(*m.IncludedSubstanceResourcesReferencedByIngredient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*m.IncludedSubstanceResourcesReferencedByIngredient))
	} else if len(*m.IncludedSubstanceResourcesReferencedByIngredient) == 1 {
		substance = &(*m.IncludedSubstanceResourcesReferencedByIngredient)[0]
	}
	return
}

func (m *MedicationPlusRelatedResources) GetIncludedOrganizationResourceReferencedByManufacturer() (organization *Organization, err error) {
	if m.IncludedOrganizationResourcesReferencedByManufacturer == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*m.IncludedOrganizationResourcesReferencedByManufacturer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*m.IncludedOrganizationResourcesReferencedByManufacturer))
	} else if len(*m.IncludedOrganizationResourcesReferencedByManufacturer) == 1 {
		organization = &(*m.IncludedOrganizationResourcesReferencedByManufacturer)[0]
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedMedicationResourcesReferencingPackageitem() (medications []Medication, err error) {
	if m.RevIncludedMedicationResourcesReferencingPackageitem == nil {
		err = errors.New("RevIncluded medications not requested")
	} else {
		medications = *m.RevIncludedMedicationResourcesReferencingPackageitem
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedMedicationResourcesReferencingIngredient() (medications []Medication, err error) {
	if m.RevIncludedMedicationResourcesReferencingIngredient == nil {
		err = errors.New("RevIncluded medications not requested")
	} else {
		medications = *m.RevIncludedMedicationResourcesReferencingIngredient
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if m.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *m.RevIncludedGroupResourcesReferencingMember
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingRequestreference
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if m.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *m.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedProtocolResourcesReferencingSubject() (protocols []Protocol, err error) {
	if m.RevIncludedProtocolResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded protocols not requested")
	} else {
		protocols = *m.RevIncludedProtocolResourcesReferencingSubject
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *m.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if m.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *m.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if m.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *m.RevIncludedListResourcesReferencingItem
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if m.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *m.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingMedication() (medicationAdministrations []MedicationAdministration, err error) {
	if m.RevIncludedMedicationAdministrationResourcesReferencingMedication == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *m.RevIncludedMedicationAdministrationResourcesReferencingMedication
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingMedication() (medicationStatements []MedicationStatement, err error) {
	if m.RevIncludedMedicationStatementResourcesReferencingMedication == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *m.RevIncludedMedicationStatementResourcesReferencingMedication
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if m.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *m.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingMedication() (medicationDispenses []MedicationDispense, err error) {
	if m.RevIncludedMedicationDispenseResourcesReferencingMedication == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *m.RevIncludedMedicationDispenseResourcesReferencingMedication
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if m.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *m.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedMedicationOrderResourcesReferencingMedication() (medicationOrders []MedicationOrder, err error) {
	if m.RevIncludedMedicationOrderResourcesReferencingMedication == nil {
		err = errors.New("RevIncluded medicationOrders not requested")
	} else {
		medicationOrders = *m.RevIncludedMedicationOrderResourcesReferencingMedication
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *m.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if m.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *m.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (m *MedicationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *m.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (m *MedicationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedMedicationResourcesReferencedByPackageitem != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByPackageitem {
			rsc := (*m.IncludedMedicationResourcesReferencedByPackageitem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByIngredient != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByIngredient {
			rsc := (*m.IncludedMedicationResourcesReferencedByIngredient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedSubstanceResourcesReferencedByIngredient != nil {
		for idx := range *m.IncludedSubstanceResourcesReferencedByIngredient {
			rsc := (*m.IncludedSubstanceResourcesReferencedByIngredient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedByManufacturer != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedByManufacturer {
			rsc := (*m.IncludedOrganizationResourcesReferencedByManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MedicationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationResourcesReferencingPackageitem != nil {
		for idx := range *m.RevIncludedMedicationResourcesReferencingPackageitem {
			rsc := (*m.RevIncludedMedicationResourcesReferencingPackageitem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationResourcesReferencingIngredient != nil {
		for idx := range *m.RevIncludedMedicationResourcesReferencingIngredient {
			rsc := (*m.RevIncludedMedicationResourcesReferencingIngredient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *m.RevIncludedGroupResourcesReferencingMember {
			rsc := (*m.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProtocolResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedProtocolResourcesReferencingSubject {
			rsc := (*m.RevIncludedProtocolResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *m.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*m.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *m.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*m.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationAdministrationResourcesReferencingMedication != nil {
		for idx := range *m.RevIncludedMedicationAdministrationResourcesReferencingMedication {
			rsc := (*m.RevIncludedMedicationAdministrationResourcesReferencingMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationStatementResourcesReferencingMedication != nil {
		for idx := range *m.RevIncludedMedicationStatementResourcesReferencingMedication {
			rsc := (*m.RevIncludedMedicationStatementResourcesReferencingMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationDispenseResourcesReferencingMedication != nil {
		for idx := range *m.RevIncludedMedicationDispenseResourcesReferencingMedication {
			rsc := (*m.RevIncludedMedicationDispenseResourcesReferencingMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationOrderResourcesReferencingMedication != nil {
		for idx := range *m.RevIncludedMedicationOrderResourcesReferencingMedication {
			rsc := (*m.RevIncludedMedicationOrderResourcesReferencingMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MedicationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedMedicationResourcesReferencedByPackageitem != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByPackageitem {
			rsc := (*m.IncludedMedicationResourcesReferencedByPackageitem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByIngredient != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByIngredient {
			rsc := (*m.IncludedMedicationResourcesReferencedByIngredient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedSubstanceResourcesReferencedByIngredient != nil {
		for idx := range *m.IncludedSubstanceResourcesReferencedByIngredient {
			rsc := (*m.IncludedSubstanceResourcesReferencedByIngredient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedByManufacturer != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedByManufacturer {
			rsc := (*m.IncludedOrganizationResourcesReferencedByManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationResourcesReferencingPackageitem != nil {
		for idx := range *m.RevIncludedMedicationResourcesReferencingPackageitem {
			rsc := (*m.RevIncludedMedicationResourcesReferencingPackageitem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationResourcesReferencingIngredient != nil {
		for idx := range *m.RevIncludedMedicationResourcesReferencingIngredient {
			rsc := (*m.RevIncludedMedicationResourcesReferencingIngredient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *m.RevIncludedGroupResourcesReferencingMember {
			rsc := (*m.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProtocolResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedProtocolResourcesReferencingSubject {
			rsc := (*m.RevIncludedProtocolResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *m.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*m.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *m.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*m.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationAdministrationResourcesReferencingMedication != nil {
		for idx := range *m.RevIncludedMedicationAdministrationResourcesReferencingMedication {
			rsc := (*m.RevIncludedMedicationAdministrationResourcesReferencingMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationStatementResourcesReferencingMedication != nil {
		for idx := range *m.RevIncludedMedicationStatementResourcesReferencingMedication {
			rsc := (*m.RevIncludedMedicationStatementResourcesReferencingMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationDispenseResourcesReferencingMedication != nil {
		for idx := range *m.RevIncludedMedicationDispenseResourcesReferencingMedication {
			rsc := (*m.RevIncludedMedicationDispenseResourcesReferencingMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationOrderResourcesReferencingMedication != nil {
		for idx := range *m.RevIncludedMedicationOrderResourcesReferencingMedication {
			rsc := (*m.RevIncludedMedicationOrderResourcesReferencingMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
