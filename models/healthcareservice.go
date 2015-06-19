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

import "time"

type HealthcareService struct {
	Id                     string                                    `json:"-" bson:"_id"`
	Identifier             []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ProvidedBy             *Reference                                `bson:"providedBy,omitempty" json:"providedBy,omitempty"`
	Location               *Reference                                `bson:"location,omitempty" json:"location,omitempty"`
	ServiceCategory        *CodeableConcept                          `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType            []HealthcareServiceServiceTypeComponent   `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	ServiceName            string                                    `bson:"serviceName,omitempty" json:"serviceName,omitempty"`
	Comment                string                                    `bson:"comment,omitempty" json:"comment,omitempty"`
	ExtraDetails           string                                    `bson:"extraDetails,omitempty" json:"extraDetails,omitempty"`
	Photo                  *Attachment                               `bson:"photo,omitempty" json:"photo,omitempty"`
	Telecom                []ContactPoint                            `bson:"telecom,omitempty" json:"telecom,omitempty"`
	CoverageArea           []Reference                               `bson:"coverageArea,omitempty" json:"coverageArea,omitempty"`
	ServiceProvisionCode   []CodeableConcept                         `bson:"serviceProvisionCode,omitempty" json:"serviceProvisionCode,omitempty"`
	Eligibility            *CodeableConcept                          `bson:"eligibility,omitempty" json:"eligibility,omitempty"`
	EligibilityNote        string                                    `bson:"eligibilityNote,omitempty" json:"eligibilityNote,omitempty"`
	ProgramName            []string                                  `bson:"programName,omitempty" json:"programName,omitempty"`
	Characteristic         []CodeableConcept                         `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	ReferralMethod         []CodeableConcept                         `bson:"referralMethod,omitempty" json:"referralMethod,omitempty"`
	PublicKey              string                                    `bson:"publicKey,omitempty" json:"publicKey,omitempty"`
	AppointmentRequired    *bool                                     `bson:"appointmentRequired,omitempty" json:"appointmentRequired,omitempty"`
	AvailableTime          []HealthcareServiceAvailableTimeComponent `bson:"availableTime,omitempty" json:"availableTime,omitempty"`
	NotAvailable           []HealthcareServiceNotAvailableComponent  `bson:"notAvailable,omitempty" json:"notAvailable,omitempty"`
	AvailabilityExceptions string                                    `bson:"availabilityExceptions,omitempty" json:"availabilityExceptions,omitempty"`
}
type HealthcareServiceServiceTypeComponent struct {
	Type      *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
}
type HealthcareServiceAvailableTimeComponent struct {
	DaysOfWeek         []string      `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty"`
	AllDay             *bool         `bson:"allDay,omitempty" json:"allDay,omitempty"`
	AvailableStartTime *FHIRDateTime `bson:"availableStartTime,omitempty" json:"availableStartTime,omitempty"`
	AvailableEndTime   *FHIRDateTime `bson:"availableEndTime,omitempty" json:"availableEndTime,omitempty"`
}
type HealthcareServiceNotAvailableComponent struct {
	Description string  `bson:"description,omitempty" json:"description,omitempty"`
	During      *Period `bson:"during,omitempty" json:"during,omitempty"`
}

type HealthcareServiceBundle struct {
	Type         string                         `json:"resourceType,omitempty"`
	Title        string                         `json:"title,omitempty"`
	Id           string                         `json:"id,omitempty"`
	Updated      time.Time                      `json:"updated,omitempty"`
	TotalResults int                            `json:"totalResults,omitempty"`
	Entry        []HealthcareServiceBundleEntry `json:"entry,omitempty"`
	Category     HealthcareServiceCategory      `json:"category,omitempty"`
}

type HealthcareServiceBundleEntry struct {
	Title    string                    `json:"title,omitempty"`
	Id       string                    `json:"id,omitempty"`
	Content  HealthcareService         `json:"content,omitempty"`
	Category HealthcareServiceCategory `json:"category,omitempty"`
}

type HealthcareServiceCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
