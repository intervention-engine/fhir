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

type Contract struct {
	DomainResource    `bson:",inline"`
	Identifier        *Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Issued            *FHIRDateTime                         `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies           *Period                               `bson:"applies,omitempty" json:"applies,omitempty"`
	Subject           []Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Authority         []Reference                           `bson:"authority,omitempty" json:"authority,omitempty"`
	Domain            []Reference                           `bson:"domain,omitempty" json:"domain,omitempty"`
	Type              *CodeableConcept                      `bson:"type,omitempty" json:"type,omitempty"`
	SubType           []CodeableConcept                     `bson:"subType,omitempty" json:"subType,omitempty"`
	Action            []CodeableConcept                     `bson:"action,omitempty" json:"action,omitempty"`
	ActionReason      []CodeableConcept                     `bson:"actionReason,omitempty" json:"actionReason,omitempty"`
	Actor             []ContractActorComponent              `bson:"actor,omitempty" json:"actor,omitempty"`
	ValuedItem        []ContractValuedItemComponent         `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
	Signer            []ContractSignatoryComponent          `bson:"signer,omitempty" json:"signer,omitempty"`
	Term              []ContractTermComponent               `bson:"term,omitempty" json:"term,omitempty"`
	BindingAttachment *Attachment                           `bson:"bindingAttachment,omitempty" json:"bindingAttachment,omitempty"`
	BindingReference  *Reference                            `bson:"bindingReference,omitempty" json:"bindingReference,omitempty"`
	Friendly          []ContractFriendlyLanguageComponent   `bson:"friendly,omitempty" json:"friendly,omitempty"`
	Legal             []ContractLegalLanguageComponent      `bson:"legal,omitempty" json:"legal,omitempty"`
	Rule              []ContractComputableLanguageComponent `bson:"rule,omitempty" json:"rule,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Contract) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Contract"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Contract), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Contract) GetBSON() (interface{}, error) {
	x.ResourceType = "Contract"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "contract" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type contract Contract

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Contract) UnmarshalJSON(data []byte) (err error) {
	x2 := contract{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Contract(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Contract) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Contract"
	} else if x.ResourceType != "Contract" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Contract, instead received %s", x.ResourceType))
	}
	return nil
}

type ContractActorComponent struct {
	BackboneElement `bson:",inline"`
	Entity          *Reference        `bson:"entity,omitempty" json:"entity,omitempty"`
	Role            []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ContractValuedItemComponent struct {
	BackboneElement       `bson:",inline"`
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty"`
	Identifier            *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime         *FHIRDateTime    `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity              *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice             *Quantity        `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                *float64         `bson:"factor,omitempty" json:"factor,omitempty"`
	Points                *float64         `bson:"points,omitempty" json:"points,omitempty"`
	Net                   *Quantity        `bson:"net,omitempty" json:"net,omitempty"`
}

type ContractSignatoryComponent struct {
	BackboneElement `bson:",inline"`
	Type            *Coding    `bson:"type,omitempty" json:"type,omitempty"`
	Party           *Reference `bson:"party,omitempty" json:"party,omitempty"`
	Signature       string     `bson:"signature,omitempty" json:"signature,omitempty"`
}

type ContractTermComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      *Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Issued          *FHIRDateTime                     `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies         *Period                           `bson:"applies,omitempty" json:"applies,omitempty"`
	Type            *CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	SubType         *CodeableConcept                  `bson:"subType,omitempty" json:"subType,omitempty"`
	Subject         *Reference                        `bson:"subject,omitempty" json:"subject,omitempty"`
	Action          []CodeableConcept                 `bson:"action,omitempty" json:"action,omitempty"`
	ActionReason    []CodeableConcept                 `bson:"actionReason,omitempty" json:"actionReason,omitempty"`
	Actor           []ContractTermActorComponent      `bson:"actor,omitempty" json:"actor,omitempty"`
	Text            string                            `bson:"text,omitempty" json:"text,omitempty"`
	ValuedItem      []ContractTermValuedItemComponent `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
	Group           []ContractTermComponent           `bson:"group,omitempty" json:"group,omitempty"`
}

type ContractTermActorComponent struct {
	BackboneElement `bson:",inline"`
	Entity          *Reference        `bson:"entity,omitempty" json:"entity,omitempty"`
	Role            []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ContractTermValuedItemComponent struct {
	BackboneElement       `bson:",inline"`
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty"`
	Identifier            *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime         *FHIRDateTime    `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity              *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice             *Quantity        `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                *float64         `bson:"factor,omitempty" json:"factor,omitempty"`
	Points                *float64         `bson:"points,omitempty" json:"points,omitempty"`
	Net                   *Quantity        `bson:"net,omitempty" json:"net,omitempty"`
}

type ContractFriendlyLanguageComponent struct {
	BackboneElement   `bson:",inline"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractLegalLanguageComponent struct {
	BackboneElement   `bson:",inline"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractComputableLanguageComponent struct {
	BackboneElement   `bson:",inline"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractPlus struct {
	Contract                     `bson:",inline"`
	ContractPlusRelatedResources `bson:",inline"`
}

type ContractPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByActor              *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByActor,omitempty"`
	IncludedGroupResourcesReferencedByActor                     *[]Group                 `bson:"_includedGroupResourcesReferencedByActor,omitempty"`
	IncludedOrganizationResourcesReferencedByActor              *[]Organization          `bson:"_includedOrganizationResourcesReferencedByActor,omitempty"`
	IncludedDeviceResourcesReferencedByActor                    *[]Device                `bson:"_includedDeviceResourcesReferencedByActor,omitempty"`
	IncludedPatientResourcesReferencedByActor                   *[]Patient               `bson:"_includedPatientResourcesReferencedByActor,omitempty"`
	IncludedSubstanceResourcesReferencedByActor                 *[]Substance             `bson:"_includedSubstanceResourcesReferencedByActor,omitempty"`
	IncludedContractResourcesReferencedByActor                  *[]Contract              `bson:"_includedContractResourcesReferencedByActor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByActor             *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByActor,omitempty"`
	IncludedLocationResourcesReferencedByActor                  *[]Location              `bson:"_includedLocationResourcesReferencedByActor,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedBySigner             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySigner,omitempty"`
	IncludedOrganizationResourcesReferencedBySigner             *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySigner,omitempty"`
	IncludedPatientResourcesReferencedBySigner                  *[]Patient               `bson:"_includedPatientResourcesReferencedBySigner,omitempty"`
	IncludedRelatedPersonResourcesReferencedBySigner            *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedBySigner,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedContractResourcesReferencingActor                *[]Contract              `bson:"_revIncludedContractResourcesReferencingActor,omitempty"`
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

func (c *ContractPlusRelatedResources) GetIncludedPractitionerResourceReferencedByActor() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByActor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByActor))
	} else if len(*c.IncludedPractitionerResourcesReferencedByActor) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByActor)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedGroupResourceReferencedByActor() (group *Group, err error) {
	if c.IncludedGroupResourcesReferencedByActor == nil {
		err = errors.New("Included groups not requested")
	} else if len(*c.IncludedGroupResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*c.IncludedGroupResourcesReferencedByActor))
	} else if len(*c.IncludedGroupResourcesReferencedByActor) == 1 {
		group = &(*c.IncludedGroupResourcesReferencedByActor)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedOrganizationResourceReferencedByActor() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByActor == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByActor))
	} else if len(*c.IncludedOrganizationResourcesReferencedByActor) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByActor)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedDeviceResourceReferencedByActor() (device *Device, err error) {
	if c.IncludedDeviceResourcesReferencedByActor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*c.IncludedDeviceResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*c.IncludedDeviceResourcesReferencedByActor))
	} else if len(*c.IncludedDeviceResourcesReferencedByActor) == 1 {
		device = &(*c.IncludedDeviceResourcesReferencedByActor)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPatientResourceReferencedByActor() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByActor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByActor))
	} else if len(*c.IncludedPatientResourcesReferencedByActor) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByActor)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedSubstanceResourceReferencedByActor() (substance *Substance, err error) {
	if c.IncludedSubstanceResourcesReferencedByActor == nil {
		err = errors.New("Included substances not requested")
	} else if len(*c.IncludedSubstanceResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*c.IncludedSubstanceResourcesReferencedByActor))
	} else if len(*c.IncludedSubstanceResourcesReferencedByActor) == 1 {
		substance = &(*c.IncludedSubstanceResourcesReferencedByActor)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedContractResourceReferencedByActor() (contract *Contract, err error) {
	if c.IncludedContractResourcesReferencedByActor == nil {
		err = errors.New("Included contracts not requested")
	} else if len(*c.IncludedContractResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 contract, but found %d", len(*c.IncludedContractResourcesReferencedByActor))
	} else if len(*c.IncludedContractResourcesReferencedByActor) == 1 {
		contract = &(*c.IncludedContractResourcesReferencedByActor)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByActor() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByActor == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedByActor))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByActor) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedByActor)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedLocationResourceReferencedByActor() (location *Location, err error) {
	if c.IncludedLocationResourcesReferencedByActor == nil {
		err = errors.New("Included locations not requested")
	} else if len(*c.IncludedLocationResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*c.IncludedLocationResourcesReferencedByActor))
	} else if len(*c.IncludedLocationResourcesReferencedByActor) == 1 {
		location = &(*c.IncludedLocationResourcesReferencedByActor)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPatientResourcesReferencedBySubject() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedBySubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPatientResourcesReferencedByPatient() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedByPatient
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySigner() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedBySigner == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedBySigner))
	} else if len(*c.IncludedPractitionerResourcesReferencedBySigner) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySigner() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedBySigner == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedBySigner))
	} else if len(*c.IncludedOrganizationResourcesReferencedBySigner) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPatientResourceReferencedBySigner() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySigner == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedBySigner))
	} else if len(*c.IncludedPatientResourcesReferencedBySigner) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedBySigner() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedBySigner == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedBySigner))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedBySigner) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedContractResourcesReferencingActor() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingActor == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingActor
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPractitionerResourcesReferencedByActor != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByActor {
			rsc := (*c.IncludedPractitionerResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedByActor != nil {
		for idx := range *c.IncludedGroupResourcesReferencedByActor {
			rsc := (*c.IncludedGroupResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByActor != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByActor {
			rsc := (*c.IncludedOrganizationResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByActor != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByActor {
			rsc := (*c.IncludedDeviceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByActor != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByActor {
			rsc := (*c.IncludedPatientResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSubstanceResourcesReferencedByActor != nil {
		for idx := range *c.IncludedSubstanceResourcesReferencedByActor {
			rsc := (*c.IncludedSubstanceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedContractResourcesReferencedByActor != nil {
		for idx := range *c.IncludedContractResourcesReferencedByActor {
			rsc := (*c.IncludedContractResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByActor != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByActor {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByActor != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByActor {
			rsc := (*c.IncludedLocationResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubject {
			rsc := (*c.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedBySigner {
			rsc := (*c.IncludedPractitionerResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedBySigner {
			rsc := (*c.IncludedOrganizationResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySigner {
			rsc := (*c.IncludedPatientResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedBySigner {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ContractPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	if c.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingActor {
			rsc := (*c.RevIncludedContractResourcesReferencingActor)[idx]
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

func (c *ContractPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPractitionerResourcesReferencedByActor != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByActor {
			rsc := (*c.IncludedPractitionerResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedByActor != nil {
		for idx := range *c.IncludedGroupResourcesReferencedByActor {
			rsc := (*c.IncludedGroupResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByActor != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByActor {
			rsc := (*c.IncludedOrganizationResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByActor != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByActor {
			rsc := (*c.IncludedDeviceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByActor != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByActor {
			rsc := (*c.IncludedPatientResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSubstanceResourcesReferencedByActor != nil {
		for idx := range *c.IncludedSubstanceResourcesReferencedByActor {
			rsc := (*c.IncludedSubstanceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedContractResourcesReferencedByActor != nil {
		for idx := range *c.IncludedContractResourcesReferencedByActor {
			rsc := (*c.IncludedContractResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByActor != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByActor {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByActor != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByActor {
			rsc := (*c.IncludedLocationResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubject {
			rsc := (*c.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedBySigner {
			rsc := (*c.IncludedPractitionerResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedBySigner {
			rsc := (*c.IncludedOrganizationResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySigner {
			rsc := (*c.IncludedPatientResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedBySigner {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedBySigner)[idx]
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
	if c.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingActor {
			rsc := (*c.RevIncludedContractResourcesReferencingActor)[idx]
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
