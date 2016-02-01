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
	Entity *Reference        `bson:"entity,omitempty" json:"entity,omitempty"`
	Role   []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ContractValuedItemComponent struct {
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
	Type      *Coding    `bson:"type,omitempty" json:"type,omitempty"`
	Party     *Reference `bson:"party,omitempty" json:"party,omitempty"`
	Signature string     `bson:"signature,omitempty" json:"signature,omitempty"`
}

type ContractTermComponent struct {
	Identifier   *Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Issued       *FHIRDateTime                     `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies      *Period                           `bson:"applies,omitempty" json:"applies,omitempty"`
	Type         *CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	SubType      *CodeableConcept                  `bson:"subType,omitempty" json:"subType,omitempty"`
	Subject      *Reference                        `bson:"subject,omitempty" json:"subject,omitempty"`
	Action       []CodeableConcept                 `bson:"action,omitempty" json:"action,omitempty"`
	ActionReason []CodeableConcept                 `bson:"actionReason,omitempty" json:"actionReason,omitempty"`
	Actor        []ContractTermActorComponent      `bson:"actor,omitempty" json:"actor,omitempty"`
	Text         string                            `bson:"text,omitempty" json:"text,omitempty"`
	ValuedItem   []ContractTermValuedItemComponent `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
	Group        []ContractTermComponent           `bson:"group,omitempty" json:"group,omitempty"`
}

type ContractTermActorComponent struct {
	Entity *Reference        `bson:"entity,omitempty" json:"entity,omitempty"`
	Role   []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ContractTermValuedItemComponent struct {
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
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractLegalLanguageComponent struct {
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractComputableLanguageComponent struct {
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractPlus struct {
	Contract             `bson:",inline"`
	ContractPlusIncludes `bson:",inline"`
}

type ContractPlusIncludes struct {
	IncludedActorPractitionerResources   *[]Practitioner  `bson:"_includedActorPractitionerResources,omitempty"`
	IncludedActorGroupResources          *[]Group         `bson:"_includedActorGroupResources,omitempty"`
	IncludedActorOrganizationResources   *[]Organization  `bson:"_includedActorOrganizationResources,omitempty"`
	IncludedActorDeviceResources         *[]Device        `bson:"_includedActorDeviceResources,omitempty"`
	IncludedActorPatientResources        *[]Patient       `bson:"_includedActorPatientResources,omitempty"`
	IncludedActorSubstanceResources      *[]Substance     `bson:"_includedActorSubstanceResources,omitempty"`
	IncludedActorContractResources       *[]Contract      `bson:"_includedActorContractResources,omitempty"`
	IncludedActorRelatedPersonResources  *[]RelatedPerson `bson:"_includedActorRelatedPersonResources,omitempty"`
	IncludedActorLocationResources       *[]Location      `bson:"_includedActorLocationResources,omitempty"`
	IncludedSubjectResources             *[]Patient       `bson:"_includedSubjectResources,omitempty"`
	IncludedPatientResources             *[]Patient       `bson:"_includedPatientResources,omitempty"`
	IncludedSignerPractitionerResources  *[]Practitioner  `bson:"_includedSignerPractitionerResources,omitempty"`
	IncludedSignerOrganizationResources  *[]Organization  `bson:"_includedSignerOrganizationResources,omitempty"`
	IncludedSignerPatientResources       *[]Patient       `bson:"_includedSignerPatientResources,omitempty"`
	IncludedSignerRelatedPersonResources *[]RelatedPerson `bson:"_includedSignerRelatedPersonResources,omitempty"`
}

func (c *ContractPlusIncludes) GetIncludedActorPractitionerResource() (practitioner *Practitioner, err error) {
	if c.IncludedActorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedActorPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedActorPractitionerResources))
	} else if len(*c.IncludedActorPractitionerResources) == 1 {
		practitioner = &(*c.IncludedActorPractitionerResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedActorGroupResource() (group *Group, err error) {
	if c.IncludedActorGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*c.IncludedActorGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*c.IncludedActorGroupResources))
	} else if len(*c.IncludedActorGroupResources) == 1 {
		group = &(*c.IncludedActorGroupResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedActorOrganizationResource() (organization *Organization, err error) {
	if c.IncludedActorOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedActorOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedActorOrganizationResources))
	} else if len(*c.IncludedActorOrganizationResources) == 1 {
		organization = &(*c.IncludedActorOrganizationResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedActorDeviceResource() (device *Device, err error) {
	if c.IncludedActorDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*c.IncludedActorDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*c.IncludedActorDeviceResources))
	} else if len(*c.IncludedActorDeviceResources) == 1 {
		device = &(*c.IncludedActorDeviceResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedActorPatientResource() (patient *Patient, err error) {
	if c.IncludedActorPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedActorPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedActorPatientResources))
	} else if len(*c.IncludedActorPatientResources) == 1 {
		patient = &(*c.IncludedActorPatientResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedActorSubstanceResource() (substance *Substance, err error) {
	if c.IncludedActorSubstanceResources == nil {
		err = errors.New("Included substances not requested")
	} else if len(*c.IncludedActorSubstanceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*c.IncludedActorSubstanceResources))
	} else if len(*c.IncludedActorSubstanceResources) == 1 {
		substance = &(*c.IncludedActorSubstanceResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedActorContractResource() (contract *Contract, err error) {
	if c.IncludedActorContractResources == nil {
		err = errors.New("Included contracts not requested")
	} else if len(*c.IncludedActorContractResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 contract, but found %d", len(*c.IncludedActorContractResources))
	} else if len(*c.IncludedActorContractResources) == 1 {
		contract = &(*c.IncludedActorContractResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedActorRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedActorRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedActorRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedActorRelatedPersonResources))
	} else if len(*c.IncludedActorRelatedPersonResources) == 1 {
		relatedPerson = &(*c.IncludedActorRelatedPersonResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedActorLocationResource() (location *Location, err error) {
	if c.IncludedActorLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*c.IncludedActorLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*c.IncludedActorLocationResources))
	} else if len(*c.IncludedActorLocationResources) == 1 {
		location = &(*c.IncludedActorLocationResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedSubjectResources() (patients []Patient, err error) {
	if c.IncludedSubjectResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedSubjectResources
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedPatientResources() (patients []Patient, err error) {
	if c.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResources
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedSignerPractitionerResource() (practitioner *Practitioner, err error) {
	if c.IncludedSignerPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedSignerPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedSignerPractitionerResources))
	} else if len(*c.IncludedSignerPractitionerResources) == 1 {
		practitioner = &(*c.IncludedSignerPractitionerResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedSignerOrganizationResource() (organization *Organization, err error) {
	if c.IncludedSignerOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedSignerOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedSignerOrganizationResources))
	} else if len(*c.IncludedSignerOrganizationResources) == 1 {
		organization = &(*c.IncludedSignerOrganizationResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedSignerPatientResource() (patient *Patient, err error) {
	if c.IncludedSignerPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedSignerPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedSignerPatientResources))
	} else if len(*c.IncludedSignerPatientResources) == 1 {
		patient = &(*c.IncludedSignerPatientResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedSignerRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedSignerRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedSignerRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedSignerRelatedPersonResources))
	} else if len(*c.IncludedSignerRelatedPersonResources) == 1 {
		relatedPerson = &(*c.IncludedSignerRelatedPersonResources)[0]
	}
	return
}

func (c *ContractPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedActorPractitionerResources != nil {
		for _, r := range *c.IncludedActorPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActorGroupResources != nil {
		for _, r := range *c.IncludedActorGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActorOrganizationResources != nil {
		for _, r := range *c.IncludedActorOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActorDeviceResources != nil {
		for _, r := range *c.IncludedActorDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActorPatientResources != nil {
		for _, r := range *c.IncludedActorPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActorSubstanceResources != nil {
		for _, r := range *c.IncludedActorSubstanceResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActorContractResources != nil {
		for _, r := range *c.IncludedActorContractResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActorRelatedPersonResources != nil {
		for _, r := range *c.IncludedActorRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActorLocationResources != nil {
		for _, r := range *c.IncludedActorLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSubjectResources != nil {
		for _, r := range *c.IncludedSubjectResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResources != nil {
		for _, r := range *c.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSignerPractitionerResources != nil {
		for _, r := range *c.IncludedSignerPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSignerOrganizationResources != nil {
		for _, r := range *c.IncludedSignerOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSignerPatientResources != nil {
		for _, r := range *c.IncludedSignerPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSignerRelatedPersonResources != nil {
		for _, r := range *c.IncludedSignerRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
