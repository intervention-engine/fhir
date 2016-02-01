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

type ImagingObjectSelection struct {
	DomainResource `bson:",inline"`
	Uid            string                                 `bson:"uid,omitempty" json:"uid,omitempty"`
	Patient        *Reference                             `bson:"patient,omitempty" json:"patient,omitempty"`
	Title          *CodeableConcept                       `bson:"title,omitempty" json:"title,omitempty"`
	Description    string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Author         *Reference                             `bson:"author,omitempty" json:"author,omitempty"`
	AuthoringTime  *FHIRDateTime                          `bson:"authoringTime,omitempty" json:"authoringTime,omitempty"`
	Study          []ImagingObjectSelectionStudyComponent `bson:"study,omitempty" json:"study,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ImagingObjectSelection) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ImagingObjectSelection"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ImagingObjectSelection), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ImagingObjectSelection) GetBSON() (interface{}, error) {
	x.ResourceType = "ImagingObjectSelection"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "imagingObjectSelection" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type imagingObjectSelection ImagingObjectSelection

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ImagingObjectSelection) UnmarshalJSON(data []byte) (err error) {
	x2 := imagingObjectSelection{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ImagingObjectSelection(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ImagingObjectSelection) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ImagingObjectSelection"
	} else if x.ResourceType != "ImagingObjectSelection" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ImagingObjectSelection, instead received %s", x.ResourceType))
	}
	return nil
}

type ImagingObjectSelectionStudyComponent struct {
	Uid          string                                  `bson:"uid,omitempty" json:"uid,omitempty"`
	Url          string                                  `bson:"url,omitempty" json:"url,omitempty"`
	ImagingStudy *Reference                              `bson:"imagingStudy,omitempty" json:"imagingStudy,omitempty"`
	Series       []ImagingObjectSelectionSeriesComponent `bson:"series,omitempty" json:"series,omitempty"`
}

type ImagingObjectSelectionSeriesComponent struct {
	Uid      string                                    `bson:"uid,omitempty" json:"uid,omitempty"`
	Url      string                                    `bson:"url,omitempty" json:"url,omitempty"`
	Instance []ImagingObjectSelectionInstanceComponent `bson:"instance,omitempty" json:"instance,omitempty"`
}

type ImagingObjectSelectionInstanceComponent struct {
	SopClass string                                  `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	Uid      string                                  `bson:"uid,omitempty" json:"uid,omitempty"`
	Url      string                                  `bson:"url,omitempty" json:"url,omitempty"`
	Frames   []ImagingObjectSelectionFramesComponent `bson:"frames,omitempty" json:"frames,omitempty"`
}

type ImagingObjectSelectionFramesComponent struct {
	FrameNumbers []uint32 `bson:"frameNumbers,omitempty" json:"frameNumbers,omitempty"`
	Url          string   `bson:"url,omitempty" json:"url,omitempty"`
}

type ImagingObjectSelectionPlus struct {
	ImagingObjectSelection             `bson:",inline"`
	ImagingObjectSelectionPlusIncludes `bson:",inline"`
}

type ImagingObjectSelectionPlusIncludes struct {
	IncludedAuthorPractitionerResources  *[]Practitioner  `bson:"_includedAuthorPractitionerResources,omitempty"`
	IncludedAuthorOrganizationResources  *[]Organization  `bson:"_includedAuthorOrganizationResources,omitempty"`
	IncludedAuthorDeviceResources        *[]Device        `bson:"_includedAuthorDeviceResources,omitempty"`
	IncludedAuthorPatientResources       *[]Patient       `bson:"_includedAuthorPatientResources,omitempty"`
	IncludedAuthorRelatedPersonResources *[]RelatedPerson `bson:"_includedAuthorRelatedPersonResources,omitempty"`
	IncludedPatientResources             *[]Patient       `bson:"_includedPatientResources,omitempty"`
}

func (i *ImagingObjectSelectionPlusIncludes) GetIncludedAuthorPractitionerResource() (practitioner *Practitioner, err error) {
	if i.IncludedAuthorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*i.IncludedAuthorPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*i.IncludedAuthorPractitionerResources))
	} else if len(*i.IncludedAuthorPractitionerResources) == 1 {
		practitioner = &(*i.IncludedAuthorPractitionerResources)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusIncludes) GetIncludedAuthorOrganizationResource() (organization *Organization, err error) {
	if i.IncludedAuthorOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*i.IncludedAuthorOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*i.IncludedAuthorOrganizationResources))
	} else if len(*i.IncludedAuthorOrganizationResources) == 1 {
		organization = &(*i.IncludedAuthorOrganizationResources)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusIncludes) GetIncludedAuthorDeviceResource() (device *Device, err error) {
	if i.IncludedAuthorDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*i.IncludedAuthorDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*i.IncludedAuthorDeviceResources))
	} else if len(*i.IncludedAuthorDeviceResources) == 1 {
		device = &(*i.IncludedAuthorDeviceResources)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusIncludes) GetIncludedAuthorPatientResource() (patient *Patient, err error) {
	if i.IncludedAuthorPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedAuthorPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedAuthorPatientResources))
	} else if len(*i.IncludedAuthorPatientResources) == 1 {
		patient = &(*i.IncludedAuthorPatientResources)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusIncludes) GetIncludedAuthorRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if i.IncludedAuthorRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*i.IncludedAuthorRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*i.IncludedAuthorRelatedPersonResources))
	} else if len(*i.IncludedAuthorRelatedPersonResources) == 1 {
		relatedPerson = &(*i.IncludedAuthorRelatedPersonResources)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if i.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResources))
	} else if len(*i.IncludedPatientResources) == 1 {
		patient = &(*i.IncludedPatientResources)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedAuthorPractitionerResources != nil {
		for _, r := range *i.IncludedAuthorPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedAuthorOrganizationResources != nil {
		for _, r := range *i.IncludedAuthorOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedAuthorDeviceResources != nil {
		for _, r := range *i.IncludedAuthorDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedAuthorPatientResources != nil {
		for _, r := range *i.IncludedAuthorPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedAuthorRelatedPersonResources != nil {
		for _, r := range *i.IncludedAuthorRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedPatientResources != nil {
		for _, r := range *i.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
