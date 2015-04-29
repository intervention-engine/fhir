// Copyright (c) 2011-2014, HL7, Inc & The MITRE Corporation
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

import "time"

type Appointment struct {
	Id             string                            `json:"-" bson:"_id"`
	Identifier     []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Priority       float64                           `bson:"priority,omitempty" json:"priority,omitempty"`
	Status         string                            `bson:"status,omitempty" json:"status,omitempty"`
	Type           *CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	Reason         *CodeableConcept                  `bson:"reason,omitempty" json:"reason,omitempty"`
	Description    string                            `bson:"description,omitempty" json:"description,omitempty"`
	Start          *FHIRDateTime                     `bson:"start,omitempty" json:"start,omitempty"`
	End            *FHIRDateTime                     `bson:"end,omitempty" json:"end,omitempty"`
	Slot           []Reference                       `bson:"slot,omitempty" json:"slot,omitempty"`
	Location       *Reference                        `bson:"location,omitempty" json:"location,omitempty"`
	Comment        string                            `bson:"comment,omitempty" json:"comment,omitempty"`
	Order          *Reference                        `bson:"order,omitempty" json:"order,omitempty"`
	Participant    []AppointmentParticipantComponent `bson:"participant,omitempty" json:"participant,omitempty"`
	LastModifiedBy *Reference                        `bson:"lastModifiedBy,omitempty" json:"lastModifiedBy,omitempty"`
	LastModified   *FHIRDateTime                     `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec participant
type AppointmentParticipantComponent struct {
	Type     []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Actor    *Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
	Required string            `bson:"required,omitempty" json:"required,omitempty"`
	Status   string            `bson:"status,omitempty" json:"status,omitempty"`
}

type AppointmentBundle struct {
	Type         string                   `json:"resourceType,omitempty"`
	Title        string                   `json:"title,omitempty"`
	Id           string                   `json:"id,omitempty"`
	Updated      time.Time                `json:"updated,omitempty"`
	TotalResults int                      `json:"totalResults,omitempty"`
	Entry        []AppointmentBundleEntry `json:"entry,omitempty"`
	Category     AppointmentCategory      `json:"category,omitempty"`
}

type AppointmentBundleEntry struct {
	Title    string              `json:"title,omitempty"`
	Id       string              `json:"id,omitempty"`
	Content  Appointment         `json:"content,omitempty"`
	Category AppointmentCategory `json:"category,omitempty"`
}

type AppointmentCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
