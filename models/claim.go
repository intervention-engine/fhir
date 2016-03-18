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

type Claim struct {
	DomainResource        `bson:",inline"`
	Type                  string                       `bson:"type,omitempty" json:"type,omitempty"`
	Identifier            []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ruleset               *Coding                      `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset       *Coding                      `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created               *FHIRDateTime                `bson:"created,omitempty" json:"created,omitempty"`
	Target                *Reference                   `bson:"target,omitempty" json:"target,omitempty"`
	Provider              *Reference                   `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization          *Reference                   `bson:"organization,omitempty" json:"organization,omitempty"`
	Use                   string                       `bson:"use,omitempty" json:"use,omitempty"`
	Priority              *Coding                      `bson:"priority,omitempty" json:"priority,omitempty"`
	FundsReserve          *Coding                      `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	Enterer               *Reference                   `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Facility              *Reference                   `bson:"facility,omitempty" json:"facility,omitempty"`
	Prescription          *Reference                   `bson:"prescription,omitempty" json:"prescription,omitempty"`
	OriginalPrescription  *Reference                   `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                 *ClaimPayeeComponent         `bson:"payee,omitempty" json:"payee,omitempty"`
	Referral              *Reference                   `bson:"referral,omitempty" json:"referral,omitempty"`
	Diagnosis             []ClaimDiagnosisComponent    `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Condition             []Coding                     `bson:"condition,omitempty" json:"condition,omitempty"`
	Patient               *Reference                   `bson:"patient,omitempty" json:"patient,omitempty"`
	Coverage              []ClaimCoverageComponent     `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Exception             []Coding                     `bson:"exception,omitempty" json:"exception,omitempty"`
	School                string                       `bson:"school,omitempty" json:"school,omitempty"`
	Accident              *FHIRDateTime                `bson:"accident,omitempty" json:"accident,omitempty"`
	AccidentType          *Coding                      `bson:"accidentType,omitempty" json:"accidentType,omitempty"`
	InterventionException []Coding                     `bson:"interventionException,omitempty" json:"interventionException,omitempty"`
	Item                  []ClaimItemsComponent        `bson:"item,omitempty" json:"item,omitempty"`
	AdditionalMaterials   []Coding                     `bson:"additionalMaterials,omitempty" json:"additionalMaterials,omitempty"`
	MissingTeeth          []ClaimMissingTeethComponent `bson:"missingTeeth,omitempty" json:"missingTeeth,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Claim) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Claim"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Claim), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Claim) GetBSON() (interface{}, error) {
	x.ResourceType = "Claim"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "claim" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type claim Claim

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Claim) UnmarshalJSON(data []byte) (err error) {
	x2 := claim{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Claim(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Claim) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Claim"
	} else if x.ResourceType != "Claim" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Claim, instead received %s", x.ResourceType))
	}
	return nil
}

type ClaimPayeeComponent struct {
	BackboneElement `bson:",inline"`
	Type            *Coding    `bson:"type,omitempty" json:"type,omitempty"`
	Provider        *Reference `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization    *Reference `bson:"organization,omitempty" json:"organization,omitempty"`
	Person          *Reference `bson:"person,omitempty" json:"person,omitempty"`
}

type ClaimDiagnosisComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32 `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Diagnosis       *Coding `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
}

type ClaimCoverageComponent struct {
	BackboneElement     `bson:",inline"`
	Sequence            *uint32    `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Focal               *bool      `bson:"focal,omitempty" json:"focal,omitempty"`
	Coverage            *Reference `bson:"coverage,omitempty" json:"coverage,omitempty"`
	BusinessArrangement string     `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	Relationship        *Coding    `bson:"relationship,omitempty" json:"relationship,omitempty"`
	PreAuthRef          []string   `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
	OriginalRuleset     *Coding    `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
}

type ClaimItemsComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32                   `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type            *Coding                   `bson:"type,omitempty" json:"type,omitempty"`
	Provider        *Reference                `bson:"provider,omitempty" json:"provider,omitempty"`
	DiagnosisLinkId []uint32                  `bson:"diagnosisLinkId,omitempty" json:"diagnosisLinkId,omitempty"`
	Service         *Coding                   `bson:"service,omitempty" json:"service,omitempty"`
	ServiceDate     *FHIRDateTime             `bson:"serviceDate,omitempty" json:"serviceDate,omitempty"`
	Quantity        *Quantity                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64                  `bson:"factor,omitempty" json:"factor,omitempty"`
	Points          *float64                  `bson:"points,omitempty" json:"points,omitempty"`
	Net             *Quantity                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi             *Coding                   `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite        *Coding                   `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite         []Coding                  `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Modifier        []Coding                  `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Detail          []ClaimDetailComponent    `bson:"detail,omitempty" json:"detail,omitempty"`
	Prosthesis      *ClaimProsthesisComponent `bson:"prosthesis,omitempty" json:"prosthesis,omitempty"`
}

type ClaimDetailComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32                   `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type            *Coding                   `bson:"type,omitempty" json:"type,omitempty"`
	Service         *Coding                   `bson:"service,omitempty" json:"service,omitempty"`
	Quantity        *Quantity                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64                  `bson:"factor,omitempty" json:"factor,omitempty"`
	Points          *float64                  `bson:"points,omitempty" json:"points,omitempty"`
	Net             *Quantity                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi             *Coding                   `bson:"udi,omitempty" json:"udi,omitempty"`
	SubDetail       []ClaimSubDetailComponent `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ClaimSubDetailComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32   `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type            *Coding   `bson:"type,omitempty" json:"type,omitempty"`
	Service         *Coding   `bson:"service,omitempty" json:"service,omitempty"`
	Quantity        *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64  `bson:"factor,omitempty" json:"factor,omitempty"`
	Points          *float64  `bson:"points,omitempty" json:"points,omitempty"`
	Net             *Quantity `bson:"net,omitempty" json:"net,omitempty"`
	Udi             *Coding   `bson:"udi,omitempty" json:"udi,omitempty"`
}

type ClaimProsthesisComponent struct {
	BackboneElement `bson:",inline"`
	Initial         *bool         `bson:"initial,omitempty" json:"initial,omitempty"`
	PriorDate       *FHIRDateTime `bson:"priorDate,omitempty" json:"priorDate,omitempty"`
	PriorMaterial   *Coding       `bson:"priorMaterial,omitempty" json:"priorMaterial,omitempty"`
}

type ClaimMissingTeethComponent struct {
	BackboneElement `bson:",inline"`
	Tooth           *Coding       `bson:"tooth,omitempty" json:"tooth,omitempty"`
	Reason          *Coding       `bson:"reason,omitempty" json:"reason,omitempty"`
	ExtractionDate  *FHIRDateTime `bson:"extractionDate,omitempty" json:"extractionDate,omitempty"`
}

type ClaimPlus struct {
	Claim                     `bson:",inline"`
	ClaimPlusRelatedResources `bson:",inline"`
}

type ClaimPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByProvider           *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByProvider,omitempty"`
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

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerResourceReferencedByProvider() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByProvider == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByProvider))
	} else if len(*c.IncludedPractitionerResourcesReferencedByProvider) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByProvider)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPatient))
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*c.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *c.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*c.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*c.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ClaimPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*c.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *c.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*c.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*c.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
