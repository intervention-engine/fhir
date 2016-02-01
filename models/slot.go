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

type Slot struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type           *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Schedule       *Reference       `bson:"schedule,omitempty" json:"schedule,omitempty"`
	FreeBusyType   string           `bson:"freeBusyType,omitempty" json:"freeBusyType,omitempty"`
	Start          *FHIRDateTime    `bson:"start,omitempty" json:"start,omitempty"`
	End            *FHIRDateTime    `bson:"end,omitempty" json:"end,omitempty"`
	Overbooked     *bool            `bson:"overbooked,omitempty" json:"overbooked,omitempty"`
	Comment        string           `bson:"comment,omitempty" json:"comment,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Slot) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Slot"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Slot), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Slot) GetBSON() (interface{}, error) {
	x.ResourceType = "Slot"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "slot" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type slot Slot

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Slot) UnmarshalJSON(data []byte) (err error) {
	x2 := slot{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Slot(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Slot) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Slot"
	} else if x.ResourceType != "Slot" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Slot, instead received %s", x.ResourceType))
	}
	return nil
}

type SlotPlus struct {
	Slot             `bson:",inline"`
	SlotPlusIncludes `bson:",inline"`
}

type SlotPlusIncludes struct {
	IncludedScheduleResources *[]Schedule `bson:"_includedScheduleResources,omitempty"`
}

func (s *SlotPlusIncludes) GetIncludedScheduleResource() (schedule *Schedule, err error) {
	if s.IncludedScheduleResources == nil {
		err = errors.New("Included schedules not requested")
	} else if len(*s.IncludedScheduleResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 schedule, but found %d", len(*s.IncludedScheduleResources))
	} else if len(*s.IncludedScheduleResources) == 1 {
		schedule = &(*s.IncludedScheduleResources)[0]
	}
	return
}

func (s *SlotPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedScheduleResources != nil {
		for _, r := range *s.IncludedScheduleResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
