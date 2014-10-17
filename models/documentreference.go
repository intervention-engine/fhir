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
	MasterIdentifier Identifier                            `bson:"masterIdentifier"`
	Identifier       []Identifier                          `bson:"identifier"`
	Subject          Reference                             `bson:"subject"`
	FhirType         CodeableConcept                       `bson:"fhirType"`
	Class            CodeableConcept                       `bson:"class"`
	Author           []Reference                           `bson:"author"`
	Custodian        Reference                             `bson:"custodian"`
	PolicyManager    string                                `bson:"policyManager"`
	Authenticator    Reference                             `bson:"authenticator"`
	Created          time.Time                             `bson:"created"`
	Indexed          time.Time                             `bson:"indexed"`
	Status           string                                `bson:"status"`
	DocStatus        CodeableConcept                       `bson:"docStatus"`
	RelatesTo        []DocumentReferenceRelatesToComponent `bson:"relatesTo"`
	Description      string                                `bson:"description"`
	Confidentiality  []CodeableConcept                     `bson:"confidentiality"`
	PrimaryLanguage  string                                `bson:"primaryLanguage"`
	MimeType         string                                `bson:"mimeType"`
	Format           string                                `bson:"format"`
	Size             float64                               `bson:"size"`
	Hash             string                                `bson:"hash"`
	Location         string                                `bson:"location"`
	Service          DocumentReferenceServiceComponent     `bson:"service"`
	Context          DocumentReferenceContextComponent     `bson:"context"`
}

// This is an ugly hack to deal with embedded structures in the spec relatesTo
type DocumentReferenceRelatesToComponent struct {
	Code   string    `bson:"code"`
	Target Reference `bson:"target"`
}

// This is an ugly hack to deal with embedded structures in the spec parameter
type DocumentReferenceServiceParameterComponent struct {
	Name  string `bson:"name"`
	Value string `bson:"value"`
}

// This is an ugly hack to deal with embedded structures in the spec service
type DocumentReferenceServiceComponent struct {
	FhirType  CodeableConcept                              `bson:"fhirType"`
	Address   string                                       `bson:"address"`
	Parameter []DocumentReferenceServiceParameterComponent `bson:"parameter"`
}

// This is an ugly hack to deal with embedded structures in the spec context
type DocumentReferenceContextComponent struct {
	Event        []CodeableConcept `bson:"event"`
	Period       Period            `bson:"period"`
	FacilityType CodeableConcept   `bson:"facilityType"`
}
