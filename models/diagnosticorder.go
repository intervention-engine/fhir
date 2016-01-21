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

type DiagnosticOrder struct {
	DomainResource        `bson:",inline"`
	Subject               *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Orderer               *Reference                      `bson:"orderer,omitempty" json:"orderer,omitempty"`
	Identifier            []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Encounter             *Reference                      `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Reason                []CodeableConcept               `bson:"reason,omitempty" json:"reason,omitempty"`
	SupportingInformation []Reference                     `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Specimen              []Reference                     `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Status                string                          `bson:"status,omitempty" json:"status,omitempty"`
	Priority              string                          `bson:"priority,omitempty" json:"priority,omitempty"`
	Event                 []DiagnosticOrderEventComponent `bson:"event,omitempty" json:"event,omitempty"`
	Item                  []DiagnosticOrderItemComponent  `bson:"item,omitempty" json:"item,omitempty"`
	Note                  []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DiagnosticOrder) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		DiagnosticOrder
	}{
		ResourceType:    "DiagnosticOrder",
		DiagnosticOrder: *resource,
	}
	return json.Marshal(x)
}

// The "diagnosticOrder" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type diagnosticOrder DiagnosticOrder

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DiagnosticOrder) UnmarshalJSON(data []byte) (err error) {
	x2 := diagnosticOrder{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DiagnosticOrder(x2)
	}
	return
}

type DiagnosticOrderEventComponent struct {
	Status      string           `bson:"status,omitempty" json:"status,omitempty"`
	Description *CodeableConcept `bson:"description,omitempty" json:"description,omitempty"`
	DateTime    *FHIRDateTime    `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Actor       *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
}

type DiagnosticOrderItemComponent struct {
	Code     *CodeableConcept                `bson:"code,omitempty" json:"code,omitempty"`
	Specimen []Reference                     `bson:"specimen,omitempty" json:"specimen,omitempty"`
	BodySite *CodeableConcept                `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Status   string                          `bson:"status,omitempty" json:"status,omitempty"`
	Event    []DiagnosticOrderEventComponent `bson:"event,omitempty" json:"event,omitempty"`
}

type DiagnosticOrderPlus struct {
	DiagnosticOrder             `bson:",inline"`
	DiagnosticOrderPlusIncludes `bson:",inline"`
}

type DiagnosticOrderPlusIncludes struct {
	IncludedSubjectGroupResources           *[]Group        `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectDeviceResources          *[]Device       `bson:"_includedSubjectDeviceResources,omitempty"`
	IncludedSubjectPatientResources         *[]Patient      `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedSubjectLocationResources        *[]Location     `bson:"_includedSubjectLocationResources,omitempty"`
	IncludedEncounterResources              *[]Encounter    `bson:"_includedEncounterResources,omitempty"`
	IncludedActorPractitionerPath1Resources *[]Practitioner `bson:"_includedActorPractitionerPath1Resources,omitempty"`
	IncludedActorPractitionerPath2Resources *[]Practitioner `bson:"_includedActorPractitionerPath2Resources,omitempty"`
	IncludedActorDevicePath1Resources       *[]Device       `bson:"_includedActorDevicePath1Resources,omitempty"`
	IncludedActorDevicePath2Resources       *[]Device       `bson:"_includedActorDevicePath2Resources,omitempty"`
	IncludedPatientResources                *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedOrdererResources                *[]Practitioner `bson:"_includedOrdererResources,omitempty"`
	IncludedSpecimenPath1Resources          *[]Specimen     `bson:"_includedSpecimenPath1Resources,omitempty"`
	IncludedSpecimenPath2Resources          *[]Specimen     `bson:"_includedSpecimenPath2Resources,omitempty"`
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if d.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedSubjectGroupResources))
	} else if len(*d.IncludedSubjectGroupResources) == 1 {
		group = &(*d.IncludedSubjectGroupResources)[0]
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedSubjectDeviceResource() (device *Device, err error) {
	if d.IncludedSubjectDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedSubjectDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedSubjectDeviceResources))
	} else if len(*d.IncludedSubjectDeviceResources) == 1 {
		device = &(*d.IncludedSubjectDeviceResources)[0]
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if d.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedSubjectPatientResources))
	} else if len(*d.IncludedSubjectPatientResources) == 1 {
		patient = &(*d.IncludedSubjectPatientResources)[0]
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedSubjectLocationResource() (location *Location, err error) {
	if d.IncludedSubjectLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*d.IncludedSubjectLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*d.IncludedSubjectLocationResources))
	} else if len(*d.IncludedSubjectLocationResources) == 1 {
		location = &(*d.IncludedSubjectLocationResources)[0]
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if d.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*d.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*d.IncludedEncounterResources))
	} else if len(*d.IncludedEncounterResources) == 1 {
		encounter = &(*d.IncludedEncounterResources)[0]
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedActorPractitionerPath1Resource() (practitioner *Practitioner, err error) {
	if d.IncludedActorPractitionerPath1Resources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedActorPractitionerPath1Resources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedActorPractitionerPath1Resources))
	} else if len(*d.IncludedActorPractitionerPath1Resources) == 1 {
		practitioner = &(*d.IncludedActorPractitionerPath1Resources)[0]
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedActorPractitionerPath2Resource() (practitioner *Practitioner, err error) {
	if d.IncludedActorPractitionerPath2Resources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedActorPractitionerPath2Resources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedActorPractitionerPath2Resources))
	} else if len(*d.IncludedActorPractitionerPath2Resources) == 1 {
		practitioner = &(*d.IncludedActorPractitionerPath2Resources)[0]
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedActorDevicePath1Resource() (device *Device, err error) {
	if d.IncludedActorDevicePath1Resources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedActorDevicePath1Resources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedActorDevicePath1Resources))
	} else if len(*d.IncludedActorDevicePath1Resources) == 1 {
		device = &(*d.IncludedActorDevicePath1Resources)[0]
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedActorDevicePath2Resource() (device *Device, err error) {
	if d.IncludedActorDevicePath2Resources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedActorDevicePath2Resources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedActorDevicePath2Resources))
	} else if len(*d.IncludedActorDevicePath2Resources) == 1 {
		device = &(*d.IncludedActorDevicePath2Resources)[0]
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if d.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResources))
	} else if len(*d.IncludedPatientResources) == 1 {
		patient = &(*d.IncludedPatientResources)[0]
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedOrdererResource() (practitioner *Practitioner, err error) {
	if d.IncludedOrdererResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedOrdererResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedOrdererResources))
	} else if len(*d.IncludedOrdererResources) == 1 {
		practitioner = &(*d.IncludedOrdererResources)[0]
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedSpecimenPath1Resources() (specimen []Specimen, err error) {
	if d.IncludedSpecimenPath1Resources == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *d.IncludedSpecimenPath1Resources
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedSpecimenPath2Resources() (specimen []Specimen, err error) {
	if d.IncludedSpecimenPath2Resources == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *d.IncludedSpecimenPath2Resources
	}
	return
}

func (d *DiagnosticOrderPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
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
	if d.IncludedSubjectLocationResources != nil {
		for _, r := range *d.IncludedSubjectLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedEncounterResources != nil {
		for _, r := range *d.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedActorPractitionerPath1Resources != nil {
		for _, r := range *d.IncludedActorPractitionerPath1Resources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedActorPractitionerPath2Resources != nil {
		for _, r := range *d.IncludedActorPractitionerPath2Resources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedActorDevicePath1Resources != nil {
		for _, r := range *d.IncludedActorDevicePath1Resources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedActorDevicePath2Resources != nil {
		for _, r := range *d.IncludedActorDevicePath2Resources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResources != nil {
		for _, r := range *d.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedOrdererResources != nil {
		for _, r := range *d.IncludedOrdererResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSpecimenPath1Resources != nil {
		for _, r := range *d.IncludedSpecimenPath1Resources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSpecimenPath2Resources != nil {
		for _, r := range *d.IncludedSpecimenPath2Resources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
