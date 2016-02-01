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

type List struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Title          string               `bson:"title,omitempty" json:"title,omitempty"`
	Code           *CodeableConcept     `bson:"code,omitempty" json:"code,omitempty"`
	Subject        *Reference           `bson:"subject,omitempty" json:"subject,omitempty"`
	Source         *Reference           `bson:"source,omitempty" json:"source,omitempty"`
	Encounter      *Reference           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Status         string               `bson:"status,omitempty" json:"status,omitempty"`
	Date           *FHIRDateTime        `bson:"date,omitempty" json:"date,omitempty"`
	OrderedBy      *CodeableConcept     `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Mode           string               `bson:"mode,omitempty" json:"mode,omitempty"`
	Note           string               `bson:"note,omitempty" json:"note,omitempty"`
	Entry          []ListEntryComponent `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason    *CodeableConcept     `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *List) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "List"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to List), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *List) GetBSON() (interface{}, error) {
	x.ResourceType = "List"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "list" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type list List

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *List) UnmarshalJSON(data []byte) (err error) {
	x2 := list{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = List(x2)
		return x.checkResourceType()
	}
	return
}

func (x *List) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "List"
	} else if x.ResourceType != "List" {
		return errors.New(fmt.Sprintf("Expected resourceType to be List, instead received %s", x.ResourceType))
	}
	return nil
}

type ListEntryComponent struct {
	Flag    *CodeableConcept `bson:"flag,omitempty" json:"flag,omitempty"`
	Deleted *bool            `bson:"deleted,omitempty" json:"deleted,omitempty"`
	Date    *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Item    *Reference       `bson:"item,omitempty" json:"item,omitempty"`
}

type ListPlus struct {
	List             `bson:",inline"`
	ListPlusIncludes `bson:",inline"`
}

type ListPlusIncludes struct {
	IncludedSubjectGroupResources       *[]Group        `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectDeviceResources      *[]Device       `bson:"_includedSubjectDeviceResources,omitempty"`
	IncludedSubjectPatientResources     *[]Patient      `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedSubjectLocationResources    *[]Location     `bson:"_includedSubjectLocationResources,omitempty"`
	IncludedPatientResources            *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedSourcePractitionerResources *[]Practitioner `bson:"_includedSourcePractitionerResources,omitempty"`
	IncludedSourceDeviceResources       *[]Device       `bson:"_includedSourceDeviceResources,omitempty"`
	IncludedSourcePatientResources      *[]Patient      `bson:"_includedSourcePatientResources,omitempty"`
	IncludedEncounterResources          *[]Encounter    `bson:"_includedEncounterResources,omitempty"`
}

func (l *ListPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if l.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*l.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*l.IncludedSubjectGroupResources))
	} else if len(*l.IncludedSubjectGroupResources) == 1 {
		group = &(*l.IncludedSubjectGroupResources)[0]
	}
	return
}

func (l *ListPlusIncludes) GetIncludedSubjectDeviceResource() (device *Device, err error) {
	if l.IncludedSubjectDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*l.IncludedSubjectDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*l.IncludedSubjectDeviceResources))
	} else if len(*l.IncludedSubjectDeviceResources) == 1 {
		device = &(*l.IncludedSubjectDeviceResources)[0]
	}
	return
}

func (l *ListPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if l.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*l.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*l.IncludedSubjectPatientResources))
	} else if len(*l.IncludedSubjectPatientResources) == 1 {
		patient = &(*l.IncludedSubjectPatientResources)[0]
	}
	return
}

func (l *ListPlusIncludes) GetIncludedSubjectLocationResource() (location *Location, err error) {
	if l.IncludedSubjectLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*l.IncludedSubjectLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*l.IncludedSubjectLocationResources))
	} else if len(*l.IncludedSubjectLocationResources) == 1 {
		location = &(*l.IncludedSubjectLocationResources)[0]
	}
	return
}

func (l *ListPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if l.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*l.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*l.IncludedPatientResources))
	} else if len(*l.IncludedPatientResources) == 1 {
		patient = &(*l.IncludedPatientResources)[0]
	}
	return
}

func (l *ListPlusIncludes) GetIncludedSourcePractitionerResource() (practitioner *Practitioner, err error) {
	if l.IncludedSourcePractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*l.IncludedSourcePractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*l.IncludedSourcePractitionerResources))
	} else if len(*l.IncludedSourcePractitionerResources) == 1 {
		practitioner = &(*l.IncludedSourcePractitionerResources)[0]
	}
	return
}

func (l *ListPlusIncludes) GetIncludedSourceDeviceResource() (device *Device, err error) {
	if l.IncludedSourceDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*l.IncludedSourceDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*l.IncludedSourceDeviceResources))
	} else if len(*l.IncludedSourceDeviceResources) == 1 {
		device = &(*l.IncludedSourceDeviceResources)[0]
	}
	return
}

func (l *ListPlusIncludes) GetIncludedSourcePatientResource() (patient *Patient, err error) {
	if l.IncludedSourcePatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*l.IncludedSourcePatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*l.IncludedSourcePatientResources))
	} else if len(*l.IncludedSourcePatientResources) == 1 {
		patient = &(*l.IncludedSourcePatientResources)[0]
	}
	return
}

func (l *ListPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if l.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*l.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*l.IncludedEncounterResources))
	} else if len(*l.IncludedEncounterResources) == 1 {
		encounter = &(*l.IncludedEncounterResources)[0]
	}
	return
}

func (l *ListPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedSubjectGroupResources != nil {
		for _, r := range *l.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedSubjectDeviceResources != nil {
		for _, r := range *l.IncludedSubjectDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedSubjectPatientResources != nil {
		for _, r := range *l.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedSubjectLocationResources != nil {
		for _, r := range *l.IncludedSubjectLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedPatientResources != nil {
		for _, r := range *l.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedSourcePractitionerResources != nil {
		for _, r := range *l.IncludedSourcePractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedSourceDeviceResources != nil {
		for _, r := range *l.IncludedSourceDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedSourcePatientResources != nil {
		for _, r := range *l.IncludedSourcePatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedEncounterResources != nil {
		for _, r := range *l.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
