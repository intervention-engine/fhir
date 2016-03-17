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

type MedicationDispense struct {
	DomainResource            `bson:",inline"`
	Identifier                *Identifier                                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                    string                                         `bson:"status,omitempty" json:"status,omitempty"`
	Patient                   *Reference                                     `bson:"patient,omitempty" json:"patient,omitempty"`
	Dispenser                 *Reference                                     `bson:"dispenser,omitempty" json:"dispenser,omitempty"`
	AuthorizingPrescription   []Reference                                    `bson:"authorizingPrescription,omitempty" json:"authorizingPrescription,omitempty"`
	Type                      *CodeableConcept                               `bson:"type,omitempty" json:"type,omitempty"`
	Quantity                  *Quantity                                      `bson:"quantity,omitempty" json:"quantity,omitempty"`
	DaysSupply                *Quantity                                      `bson:"daysSupply,omitempty" json:"daysSupply,omitempty"`
	MedicationCodeableConcept *CodeableConcept                               `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                                     `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	WhenPrepared              *FHIRDateTime                                  `bson:"whenPrepared,omitempty" json:"whenPrepared,omitempty"`
	WhenHandedOver            *FHIRDateTime                                  `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty"`
	Destination               *Reference                                     `bson:"destination,omitempty" json:"destination,omitempty"`
	Receiver                  []Reference                                    `bson:"receiver,omitempty" json:"receiver,omitempty"`
	Note                      string                                         `bson:"note,omitempty" json:"note,omitempty"`
	DosageInstruction         []MedicationDispenseDosageInstructionComponent `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	Substitution              *MedicationDispenseSubstitutionComponent       `bson:"substitution,omitempty" json:"substitution,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *MedicationDispense) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "MedicationDispense"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to MedicationDispense), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *MedicationDispense) GetBSON() (interface{}, error) {
	x.ResourceType = "MedicationDispense"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "medicationDispense" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medicationDispense MedicationDispense

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *MedicationDispense) UnmarshalJSON(data []byte) (err error) {
	x2 := medicationDispense{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MedicationDispense(x2)
		return x.checkResourceType()
	}
	return
}

func (x *MedicationDispense) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "MedicationDispense"
	} else if x.ResourceType != "MedicationDispense" {
		return errors.New(fmt.Sprintf("Expected resourceType to be MedicationDispense, instead received %s", x.ResourceType))
	}
	return nil
}

type MedicationDispenseDosageInstructionComponent struct {
	BackboneElement         `bson:",inline"`
	Text                    string           `bson:"text,omitempty" json:"text,omitempty"`
	AdditionalInstructions  *CodeableConcept `bson:"additionalInstructions,omitempty" json:"additionalInstructions,omitempty"`
	Timing                  *Timing          `bson:"timing,omitempty" json:"timing,omitempty"`
	AsNeededBoolean         *bool            `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	SiteCodeableConcept     *CodeableConcept `bson:"siteCodeableConcept,omitempty" json:"siteCodeableConcept,omitempty"`
	SiteReference           *Reference       `bson:"siteReference,omitempty" json:"siteReference,omitempty"`
	Route                   *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	Method                  *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	DoseRange               *Range           `bson:"doseRange,omitempty" json:"doseRange,omitempty"`
	DoseSimpleQuantity      *Quantity        `bson:"doseSimpleQuantity,omitempty" json:"doseSimpleQuantity,omitempty"`
	RateRatio               *Ratio           `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateRange               *Range           `bson:"rateRange,omitempty" json:"rateRange,omitempty"`
	MaxDosePerPeriod        *Ratio           `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty"`
}

type MedicationDispenseSubstitutionComponent struct {
	BackboneElement  `bson:",inline"`
	Type             *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Reason           []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	ResponsibleParty []Reference       `bson:"responsibleParty,omitempty" json:"responsibleParty,omitempty"`
}

type MedicationDispensePlus struct {
	MedicationDispense                     `bson:",inline"`
	MedicationDispensePlusRelatedResources `bson:",inline"`
}

type MedicationDispensePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByReceiver           *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByReceiver,omitempty"`
	IncludedPatientResourcesReferencedByReceiver                *[]Patient               `bson:"_includedPatientResourcesReferencedByReceiver,omitempty"`
	IncludedLocationResourcesReferencedByDestination            *[]Location              `bson:"_includedLocationResourcesReferencedByDestination,omitempty"`
	IncludedMedicationResourcesReferencedByMedication           *[]Medication            `bson:"_includedMedicationResourcesReferencedByMedication,omitempty"`
	IncludedPractitionerResourcesReferencedByResponsibleparty   *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByResponsibleparty,omitempty"`
	IncludedPractitionerResourcesReferencedByDispenser          *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByDispenser,omitempty"`
	IncludedMedicationOrderResourcesReferencedByPrescription    *[]MedicationOrder       `bson:"_includedMedicationOrderResourcesReferencedByPrescription,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
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
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedPractitionerResourcesReferencedByReceiver() (practitioners []Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByReceiver == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *m.IncludedPractitionerResourcesReferencedByReceiver
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedPatientResourcesReferencedByReceiver() (patients []Patient, err error) {
	if m.IncludedPatientResourcesReferencedByReceiver == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *m.IncludedPatientResourcesReferencedByReceiver
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedLocationResourceReferencedByDestination() (location *Location, err error) {
	if m.IncludedLocationResourcesReferencedByDestination == nil {
		err = errors.New("Included locations not requested")
	} else if len(*m.IncludedLocationResourcesReferencedByDestination) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*m.IncludedLocationResourcesReferencedByDestination))
	} else if len(*m.IncludedLocationResourcesReferencedByDestination) == 1 {
		location = &(*m.IncludedLocationResourcesReferencedByDestination)[0]
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedMedicationResourceReferencedByMedication() (medication *Medication, err error) {
	if m.IncludedMedicationResourcesReferencedByMedication == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedMedicationResourcesReferencedByMedication) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedMedicationResourcesReferencedByMedication))
	} else if len(*m.IncludedMedicationResourcesReferencedByMedication) == 1 {
		medication = &(*m.IncludedMedicationResourcesReferencedByMedication)[0]
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedPractitionerResourcesReferencedByResponsibleparty() (practitioners []Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByResponsibleparty == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *m.IncludedPractitionerResourcesReferencedByResponsibleparty
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedPractitionerResourceReferencedByDispenser() (practitioner *Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByDispenser == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedPractitionerResourcesReferencedByDispenser) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedPractitionerResourcesReferencedByDispenser))
	} else if len(*m.IncludedPractitionerResourcesReferencedByDispenser) == 1 {
		practitioner = &(*m.IncludedPractitionerResourcesReferencedByDispenser)[0]
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedMedicationOrderResourcesReferencedByPrescription() (medicationOrders []MedicationOrder, err error) {
	if m.IncludedMedicationOrderResourcesReferencedByPrescription == nil {
		err = errors.New("Included medicationOrders not requested")
	} else {
		medicationOrders = *m.IncludedMedicationOrderResourcesReferencedByPrescription
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if m.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResourcesReferencedByPatient))
	} else if len(*m.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*m.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if m.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *m.RevIncludedListResourcesReferencingItem
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if m.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *m.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if m.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *m.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if m.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *m.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *m.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *m.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if m.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *m.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *m.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if m.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *m.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedPractitionerResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByReceiver {
			rsc := (*m.IncludedPractitionerResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByReceiver {
			rsc := (*m.IncludedPatientResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedLocationResourcesReferencedByDestination != nil {
		for idx := range *m.IncludedLocationResourcesReferencedByDestination {
			rsc := (*m.IncludedLocationResourcesReferencedByDestination)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByMedication != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByMedication {
			rsc := (*m.IncludedMedicationResourcesReferencedByMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByResponsibleparty != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByResponsibleparty {
			rsc := (*m.IncludedPractitionerResourcesReferencedByResponsibleparty)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByDispenser != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByDispenser {
			rsc := (*m.IncludedPractitionerResourcesReferencedByDispenser)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationOrderResourcesReferencedByPrescription != nil {
		for idx := range *m.IncludedMedicationOrderResourcesReferencedByPrescription {
			rsc := (*m.IncludedMedicationOrderResourcesReferencedByPrescription)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPatient {
			rsc := (*m.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MedicationDispensePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *m.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*m.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *m.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*m.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MedicationDispensePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedPractitionerResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByReceiver {
			rsc := (*m.IncludedPractitionerResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByReceiver {
			rsc := (*m.IncludedPatientResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedLocationResourcesReferencedByDestination != nil {
		for idx := range *m.IncludedLocationResourcesReferencedByDestination {
			rsc := (*m.IncludedLocationResourcesReferencedByDestination)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByMedication != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByMedication {
			rsc := (*m.IncludedMedicationResourcesReferencedByMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByResponsibleparty != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByResponsibleparty {
			rsc := (*m.IncludedPractitionerResourcesReferencedByResponsibleparty)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByDispenser != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByDispenser {
			rsc := (*m.IncludedPractitionerResourcesReferencedByDispenser)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationOrderResourcesReferencedByPrescription != nil {
		for idx := range *m.IncludedMedicationOrderResourcesReferencedByPrescription {
			rsc := (*m.IncludedMedicationOrderResourcesReferencedByPrescription)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPatient {
			rsc := (*m.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *m.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*m.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *m.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*m.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
