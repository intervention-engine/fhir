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

type DocumentReference struct {
	Id               string                                `json:"-" bson:"_id"`
	MasterIdentifier *Identifier                           `bson:"masterIdentifier,omitempty", json:"masterIdentifier,omitempty"`
	Identifier       []Identifier                          `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Subject          *Reference                            `bson:"subject,omitempty", json:"subject,omitempty"`
	Type             *CodeableConcept                      `bson:"type,omitempty", json:"type,omitempty"`
	Class            *CodeableConcept                      `bson:"class,omitempty", json:"class,omitempty"`
	Author           []Reference                           `bson:"author,omitempty", json:"author,omitempty"`
	Custodian        *Reference                            `bson:"custodian,omitempty", json:"custodian,omitempty"`
	PolicyManager    string                                `bson:"policyManager,omitempty", json:"policyManager,omitempty"`
	Authenticator    *Reference                            `bson:"authenticator,omitempty", json:"authenticator,omitempty"`
	Created          *FHIRDateTime                         `bson:"created,omitempty", json:"created,omitempty"`
	Indexed          *FHIRDateTime                         `bson:"indexed,omitempty", json:"indexed,omitempty"`
	Status           string                                `bson:"status,omitempty", json:"status,omitempty"`
	DocStatus        *CodeableConcept                      `bson:"docStatus,omitempty", json:"docStatus,omitempty"`
	RelatesTo        []DocumentReferenceRelatesToComponent `bson:"relatesTo,omitempty", json:"relatesTo,omitempty"`
	Description      string                                `bson:"description,omitempty", json:"description,omitempty"`
	Confidentiality  []CodeableConcept                     `bson:"confidentiality,omitempty", json:"confidentiality,omitempty"`
	PrimaryLanguage  string                                `bson:"primaryLanguage,omitempty", json:"primaryLanguage,omitempty"`
	MimeType         string                                `bson:"mimeType,omitempty", json:"mimeType,omitempty"`
	Format           []string                              `bson:"format,omitempty", json:"format,omitempty"`
	Size             float64                               `bson:"size,omitempty", json:"size,omitempty"`
	Hash             string                                `bson:"hash,omitempty", json:"hash,omitempty"`
	Location         string                                `bson:"location,omitempty", json:"location,omitempty"`
	Service          *DocumentReferenceServiceComponent    `bson:"service,omitempty", json:"service,omitempty"`
	Context          *DocumentReferenceContextComponent    `bson:"context,omitempty", json:"context,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec relatesTo
type DocumentReferenceRelatesToComponent struct {
	Code   string     `bson:"code,omitempty", json:"code,omitempty"`
	Target *Reference `bson:"target,omitempty", json:"target,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec parameter
type DocumentReferenceServiceParameterComponent struct {
	Name  string `bson:"name,omitempty", json:"name,omitempty"`
	Value string `bson:"value,omitempty", json:"value,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec service
type DocumentReferenceServiceComponent struct {
	Type      *CodeableConcept                             `bson:"type,omitempty", json:"type,omitempty"`
	Address   string                                       `bson:"address,omitempty", json:"address,omitempty"`
	Parameter []DocumentReferenceServiceParameterComponent `bson:"parameter,omitempty", json:"parameter,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec context
type DocumentReferenceContextComponent struct {
	Event        []CodeableConcept `bson:"event,omitempty", json:"event,omitempty"`
	Period       *Period           `bson:"period,omitempty", json:"period,omitempty"`
	FacilityType *CodeableConcept  `bson:"facilityType,omitempty", json:"facilityType,omitempty"`
}

type DocumentReferenceBundle struct {
	Type         string                         `json:"resourceType,omitempty"`
	Title        string                         `json:"title,omitempty"`
	Id           string                         `json:"id,omitempty"`
	Updated      time.Time                      `json:"updated,omitempty"`
	TotalResults int                            `json:"totalResults,omitempty"`
	Entry        []DocumentReferenceBundleEntry `json:"entry,omitempty"`
	Category     DocumentReferenceCategory      `json:"category,omitempty"`
}

type DocumentReferenceBundleEntry struct {
	Title    string                    `json:"title,omitempty"`
	Id       string                    `json:"id,omitempty"`
	Content  DocumentReference         `json:"content,omitempty"`
	Category DocumentReferenceCategory `json:"category,omitempty"`
}

type DocumentReferenceCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
