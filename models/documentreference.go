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

type DocumentReference struct {
	DomainResource   `bson:",inline"`
	MasterIdentifier *Identifier                           `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"`
	Identifier       []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject          *Reference                            `bson:"subject,omitempty" json:"subject,omitempty"`
	Type             *CodeableConcept                      `bson:"type,omitempty" json:"type,omitempty"`
	Class            *CodeableConcept                      `bson:"class,omitempty" json:"class,omitempty"`
	Author           []Reference                           `bson:"author,omitempty" json:"author,omitempty"`
	Custodian        *Reference                            `bson:"custodian,omitempty" json:"custodian,omitempty"`
	Authenticator    *Reference                            `bson:"authenticator,omitempty" json:"authenticator,omitempty"`
	Created          *FHIRDateTime                         `bson:"created,omitempty" json:"created,omitempty"`
	Indexed          *FHIRDateTime                         `bson:"indexed,omitempty" json:"indexed,omitempty"`
	Status           string                                `bson:"status,omitempty" json:"status,omitempty"`
	DocStatus        *CodeableConcept                      `bson:"docStatus,omitempty" json:"docStatus,omitempty"`
	RelatesTo        []DocumentReferenceRelatesToComponent `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Description      string                                `bson:"description,omitempty" json:"description,omitempty"`
	SecurityLabel    []CodeableConcept                     `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Content          []DocumentReferenceContentComponent   `bson:"content,omitempty" json:"content,omitempty"`
	Context          *DocumentReferenceContextComponent    `bson:"context,omitempty" json:"context,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DocumentReference) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DocumentReference"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DocumentReference), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DocumentReference) GetBSON() (interface{}, error) {
	x.ResourceType = "DocumentReference"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "documentReference" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type documentReference DocumentReference

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DocumentReference) UnmarshalJSON(data []byte) (err error) {
	x2 := documentReference{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DocumentReference(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DocumentReference) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DocumentReference"
	} else if x.ResourceType != "DocumentReference" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DocumentReference, instead received %s", x.ResourceType))
	}
	return nil
}

type DocumentReferenceRelatesToComponent struct {
	Code   string     `bson:"code,omitempty" json:"code,omitempty"`
	Target *Reference `bson:"target,omitempty" json:"target,omitempty"`
}

type DocumentReferenceContentComponent struct {
	Attachment *Attachment `bson:"attachment,omitempty" json:"attachment,omitempty"`
	Format     []Coding    `bson:"format,omitempty" json:"format,omitempty"`
}

type DocumentReferenceContextComponent struct {
	Encounter         *Reference                                 `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Event             []CodeableConcept                          `bson:"event,omitempty" json:"event,omitempty"`
	Period            *Period                                    `bson:"period,omitempty" json:"period,omitempty"`
	FacilityType      *CodeableConcept                           `bson:"facilityType,omitempty" json:"facilityType,omitempty"`
	PracticeSetting   *CodeableConcept                           `bson:"practiceSetting,omitempty" json:"practiceSetting,omitempty"`
	SourcePatientInfo *Reference                                 `bson:"sourcePatientInfo,omitempty" json:"sourcePatientInfo,omitempty"`
	Related           []DocumentReferenceContextRelatedComponent `bson:"related,omitempty" json:"related,omitempty"`
}

type DocumentReferenceContextRelatedComponent struct {
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ref        *Reference  `bson:"ref,omitempty" json:"ref,omitempty"`
}

type DocumentReferencePlus struct {
	DocumentReference             `bson:",inline"`
	DocumentReferencePlusIncludes `bson:",inline"`
}

type DocumentReferencePlusIncludes struct {
	IncludedSubjectPractitionerResources       *[]Practitioner      `bson:"_includedSubjectPractitionerResources,omitempty"`
	IncludedSubjectGroupResources              *[]Group             `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectDeviceResources             *[]Device            `bson:"_includedSubjectDeviceResources,omitempty"`
	IncludedSubjectPatientResources            *[]Patient           `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedPatientResources                   *[]Patient           `bson:"_includedPatientResources,omitempty"`
	IncludedAuthenticatorPractitionerResources *[]Practitioner      `bson:"_includedAuthenticatorPractitionerResources,omitempty"`
	IncludedAuthenticatorOrganizationResources *[]Organization      `bson:"_includedAuthenticatorOrganizationResources,omitempty"`
	IncludedCustodianResources                 *[]Organization      `bson:"_includedCustodianResources,omitempty"`
	IncludedAuthorPractitionerResources        *[]Practitioner      `bson:"_includedAuthorPractitionerResources,omitempty"`
	IncludedAuthorOrganizationResources        *[]Organization      `bson:"_includedAuthorOrganizationResources,omitempty"`
	IncludedAuthorDeviceResources              *[]Device            `bson:"_includedAuthorDeviceResources,omitempty"`
	IncludedAuthorPatientResources             *[]Patient           `bson:"_includedAuthorPatientResources,omitempty"`
	IncludedAuthorRelatedPersonResources       *[]RelatedPerson     `bson:"_includedAuthorRelatedPersonResources,omitempty"`
	IncludedEncounterResources                 *[]Encounter         `bson:"_includedEncounterResources,omitempty"`
	IncludedRelatestoResources                 *[]DocumentReference `bson:"_includedRelatestoResources,omitempty"`
}

func (d *DocumentReferencePlusIncludes) GetIncludedSubjectPractitionerResource() (practitioner *Practitioner, err error) {
	if d.IncludedSubjectPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedSubjectPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedSubjectPractitionerResources))
	} else if len(*d.IncludedSubjectPractitionerResources) == 1 {
		practitioner = &(*d.IncludedSubjectPractitionerResources)[0]
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if d.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedSubjectGroupResources))
	} else if len(*d.IncludedSubjectGroupResources) == 1 {
		group = &(*d.IncludedSubjectGroupResources)[0]
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedSubjectDeviceResource() (device *Device, err error) {
	if d.IncludedSubjectDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedSubjectDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedSubjectDeviceResources))
	} else if len(*d.IncludedSubjectDeviceResources) == 1 {
		device = &(*d.IncludedSubjectDeviceResources)[0]
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if d.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedSubjectPatientResources))
	} else if len(*d.IncludedSubjectPatientResources) == 1 {
		patient = &(*d.IncludedSubjectPatientResources)[0]
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if d.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResources))
	} else if len(*d.IncludedPatientResources) == 1 {
		patient = &(*d.IncludedPatientResources)[0]
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedAuthenticatorPractitionerResource() (practitioner *Practitioner, err error) {
	if d.IncludedAuthenticatorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedAuthenticatorPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedAuthenticatorPractitionerResources))
	} else if len(*d.IncludedAuthenticatorPractitionerResources) == 1 {
		practitioner = &(*d.IncludedAuthenticatorPractitionerResources)[0]
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedAuthenticatorOrganizationResource() (organization *Organization, err error) {
	if d.IncludedAuthenticatorOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*d.IncludedAuthenticatorOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*d.IncludedAuthenticatorOrganizationResources))
	} else if len(*d.IncludedAuthenticatorOrganizationResources) == 1 {
		organization = &(*d.IncludedAuthenticatorOrganizationResources)[0]
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedCustodianResource() (organization *Organization, err error) {
	if d.IncludedCustodianResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*d.IncludedCustodianResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*d.IncludedCustodianResources))
	} else if len(*d.IncludedCustodianResources) == 1 {
		organization = &(*d.IncludedCustodianResources)[0]
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedAuthorPractitionerResources() (practitioners []Practitioner, err error) {
	if d.IncludedAuthorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *d.IncludedAuthorPractitionerResources
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedAuthorOrganizationResources() (organizations []Organization, err error) {
	if d.IncludedAuthorOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *d.IncludedAuthorOrganizationResources
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedAuthorDeviceResources() (devices []Device, err error) {
	if d.IncludedAuthorDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *d.IncludedAuthorDeviceResources
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedAuthorPatientResources() (patients []Patient, err error) {
	if d.IncludedAuthorPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *d.IncludedAuthorPatientResources
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedAuthorRelatedPersonResources() (relatedPeople []RelatedPerson, err error) {
	if d.IncludedAuthorRelatedPersonResources == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *d.IncludedAuthorRelatedPersonResources
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if d.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*d.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*d.IncludedEncounterResources))
	} else if len(*d.IncludedEncounterResources) == 1 {
		encounter = &(*d.IncludedEncounterResources)[0]
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedRelatestoResource() (documentReference *DocumentReference, err error) {
	if d.IncludedRelatestoResources == nil {
		err = errors.New("Included documentreferences not requested")
	} else if len(*d.IncludedRelatestoResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 documentReference, but found %d", len(*d.IncludedRelatestoResources))
	} else if len(*d.IncludedRelatestoResources) == 1 {
		documentReference = &(*d.IncludedRelatestoResources)[0]
	}
	return
}

func (d *DocumentReferencePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedSubjectPractitionerResources != nil {
		for _, r := range *d.IncludedSubjectPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSubjectGroupResources != nil {
		for _, r := range *d.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSubjectDeviceResources != nil {
		for _, r := range *d.IncludedSubjectDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSubjectPatientResources != nil {
		for _, r := range *d.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResources != nil {
		for _, r := range *d.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedAuthenticatorPractitionerResources != nil {
		for _, r := range *d.IncludedAuthenticatorPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedAuthenticatorOrganizationResources != nil {
		for _, r := range *d.IncludedAuthenticatorOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedCustodianResources != nil {
		for _, r := range *d.IncludedCustodianResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedAuthorPractitionerResources != nil {
		for _, r := range *d.IncludedAuthorPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedAuthorOrganizationResources != nil {
		for _, r := range *d.IncludedAuthorOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedAuthorDeviceResources != nil {
		for _, r := range *d.IncludedAuthorDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedAuthorPatientResources != nil {
		for _, r := range *d.IncludedAuthorPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedAuthorRelatedPersonResources != nil {
		for _, r := range *d.IncludedAuthorRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedEncounterResources != nil {
		for _, r := range *d.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRelatestoResources != nil {
		for _, r := range *d.IncludedRelatestoResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
