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

type DocumentManifest struct {
	DomainResource   `bson:",inline"`
	MasterIdentifier *Identifier                        `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"`
	Identifier       []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject          *Reference                         `bson:"subject,omitempty" json:"subject,omitempty"`
	Recipient        []Reference                        `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Type             *CodeableConcept                   `bson:"type,omitempty" json:"type,omitempty"`
	Author           []Reference                        `bson:"author,omitempty" json:"author,omitempty"`
	Created          *FHIRDateTime                      `bson:"created,omitempty" json:"created,omitempty"`
	Source           string                             `bson:"source,omitempty" json:"source,omitempty"`
	Status           string                             `bson:"status,omitempty" json:"status,omitempty"`
	Description      string                             `bson:"description,omitempty" json:"description,omitempty"`
	Content          []DocumentManifestContentComponent `bson:"content,omitempty" json:"content,omitempty"`
	Related          []DocumentManifestRelatedComponent `bson:"related,omitempty" json:"related,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DocumentManifest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DocumentManifest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DocumentManifest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DocumentManifest) GetBSON() (interface{}, error) {
	x.ResourceType = "DocumentManifest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "documentManifest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type documentManifest DocumentManifest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DocumentManifest) UnmarshalJSON(data []byte) (err error) {
	x2 := documentManifest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DocumentManifest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DocumentManifest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DocumentManifest"
	} else if x.ResourceType != "DocumentManifest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DocumentManifest, instead received %s", x.ResourceType))
	}
	return nil
}

type DocumentManifestContentComponent struct {
	PAttachment *Attachment `bson:"pAttachment,omitempty" json:"pAttachment,omitempty"`
	PReference  *Reference  `bson:"pReference,omitempty" json:"pReference,omitempty"`
}

type DocumentManifestRelatedComponent struct {
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ref        *Reference  `bson:"ref,omitempty" json:"ref,omitempty"`
}

type DocumentManifestPlus struct {
	DocumentManifest             `bson:",inline"`
	DocumentManifestPlusIncludes `bson:",inline"`
}

type DocumentManifestPlusIncludes struct {
	IncludedSubjectPractitionerResources    *[]Practitioner  `bson:"_includedSubjectPractitionerResources,omitempty"`
	IncludedSubjectGroupResources           *[]Group         `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectDeviceResources          *[]Device        `bson:"_includedSubjectDeviceResources,omitempty"`
	IncludedSubjectPatientResources         *[]Patient       `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedAuthorPractitionerResources     *[]Practitioner  `bson:"_includedAuthorPractitionerResources,omitempty"`
	IncludedAuthorOrganizationResources     *[]Organization  `bson:"_includedAuthorOrganizationResources,omitempty"`
	IncludedAuthorDeviceResources           *[]Device        `bson:"_includedAuthorDeviceResources,omitempty"`
	IncludedAuthorPatientResources          *[]Patient       `bson:"_includedAuthorPatientResources,omitempty"`
	IncludedAuthorRelatedPersonResources    *[]RelatedPerson `bson:"_includedAuthorRelatedPersonResources,omitempty"`
	IncludedPatientResources                *[]Patient       `bson:"_includedPatientResources,omitempty"`
	IncludedRecipientPractitionerResources  *[]Practitioner  `bson:"_includedRecipientPractitionerResources,omitempty"`
	IncludedRecipientOrganizationResources  *[]Organization  `bson:"_includedRecipientOrganizationResources,omitempty"`
	IncludedRecipientPatientResources       *[]Patient       `bson:"_includedRecipientPatientResources,omitempty"`
	IncludedRecipientRelatedPersonResources *[]RelatedPerson `bson:"_includedRecipientRelatedPersonResources,omitempty"`
}

func (d *DocumentManifestPlusIncludes) GetIncludedSubjectPractitionerResource() (practitioner *Practitioner, err error) {
	if d.IncludedSubjectPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedSubjectPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedSubjectPractitionerResources))
	} else if len(*d.IncludedSubjectPractitionerResources) == 1 {
		practitioner = &(*d.IncludedSubjectPractitionerResources)[0]
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if d.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedSubjectGroupResources))
	} else if len(*d.IncludedSubjectGroupResources) == 1 {
		group = &(*d.IncludedSubjectGroupResources)[0]
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedSubjectDeviceResource() (device *Device, err error) {
	if d.IncludedSubjectDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedSubjectDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedSubjectDeviceResources))
	} else if len(*d.IncludedSubjectDeviceResources) == 1 {
		device = &(*d.IncludedSubjectDeviceResources)[0]
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if d.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedSubjectPatientResources))
	} else if len(*d.IncludedSubjectPatientResources) == 1 {
		patient = &(*d.IncludedSubjectPatientResources)[0]
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedAuthorPractitionerResources() (practitioners []Practitioner, err error) {
	if d.IncludedAuthorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *d.IncludedAuthorPractitionerResources
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedAuthorOrganizationResources() (organizations []Organization, err error) {
	if d.IncludedAuthorOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *d.IncludedAuthorOrganizationResources
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedAuthorDeviceResources() (devices []Device, err error) {
	if d.IncludedAuthorDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *d.IncludedAuthorDeviceResources
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedAuthorPatientResources() (patients []Patient, err error) {
	if d.IncludedAuthorPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *d.IncludedAuthorPatientResources
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedAuthorRelatedPersonResources() (relatedPeople []RelatedPerson, err error) {
	if d.IncludedAuthorRelatedPersonResources == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *d.IncludedAuthorRelatedPersonResources
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if d.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResources))
	} else if len(*d.IncludedPatientResources) == 1 {
		patient = &(*d.IncludedPatientResources)[0]
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedRecipientPractitionerResources() (practitioners []Practitioner, err error) {
	if d.IncludedRecipientPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *d.IncludedRecipientPractitionerResources
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedRecipientOrganizationResources() (organizations []Organization, err error) {
	if d.IncludedRecipientOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *d.IncludedRecipientOrganizationResources
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedRecipientPatientResources() (patients []Patient, err error) {
	if d.IncludedRecipientPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *d.IncludedRecipientPatientResources
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedRecipientRelatedPersonResources() (relatedPeople []RelatedPerson, err error) {
	if d.IncludedRecipientRelatedPersonResources == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *d.IncludedRecipientRelatedPersonResources
	}
	return
}

func (d *DocumentManifestPlusIncludes) GetIncludedResources() map[string]interface{} {
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
	if d.IncludedPatientResources != nil {
		for _, r := range *d.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRecipientPractitionerResources != nil {
		for _, r := range *d.IncludedRecipientPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRecipientOrganizationResources != nil {
		for _, r := range *d.IncludedRecipientOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRecipientPatientResources != nil {
		for _, r := range *d.IncludedRecipientPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRecipientRelatedPersonResources != nil {
		for _, r := range *d.IncludedRecipientRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
