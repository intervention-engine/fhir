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

type Conformance struct {
	Id             string                             `json:"-" bson:"_id"`
	Identifier     string                             `bson:"identifier"`
	Version        string                             `bson:"version"`
	Name           string                             `bson:"name"`
	Publisher      string                             `bson:"publisher"`
	Telecom        []ContactPoint                     `bson:"telecom"`
	Description    string                             `bson:"description"`
	Status         string                             `bson:"status"`
	Experimental   bool                               `bson:"experimental"`
	Date           time.Time                          `bson:"date"`
	Software       ConformanceSoftwareComponent       `bson:"software"`
	Implementation ConformanceImplementationComponent `bson:"implementation"`
	FhirVersion    string                             `bson:"fhirVersion"`
	AcceptUnknown  bool                               `bson:"acceptUnknown"`
	Format         []string                           `bson:"format"`
	Profile        []Reference                        `bson:"profile"`
	Rest           []ConformanceRestComponent         `bson:"rest"`
	Messaging      []ConformanceMessagingComponent    `bson:"messaging"`
	Document       []ConformanceDocumentComponent     `bson:"document"`
}

// This is an ugly hack to deal with embedded structures in the spec software
type ConformanceSoftwareComponent struct {
	Name        string    `bson:"name"`
	Version     string    `bson:"version"`
	ReleaseDate time.Time `bson:"releaseDate"`
}

// This is an ugly hack to deal with embedded structures in the spec implementation
type ConformanceImplementationComponent struct {
	Description string `bson:"description"`
	Url         string `bson:"url"`
}

// This is an ugly hack to deal with embedded structures in the spec certificate
type ConformanceRestSecurityCertificateComponent struct {
	Type string `bson:"type"`
}

// This is an ugly hack to deal with embedded structures in the spec security
type ConformanceRestSecurityComponent struct {
	Cors        bool                                          `bson:"cors"`
	Service     []CodeableConcept                             `bson:"service"`
	Description string                                        `bson:"description"`
	Certificate []ConformanceRestSecurityCertificateComponent `bson:"certificate"`
}

// This is an ugly hack to deal with embedded structures in the spec interaction
type ResourceInteractionComponent struct {
	Code          string `bson:"code"`
	Documentation string `bson:"documentation"`
}

// This is an ugly hack to deal with embedded structures in the spec searchParam
type ConformanceRestResourceSearchParamComponent struct {
	Name          string   `bson:"name"`
	Definition    string   `bson:"definition"`
	Type          string   `bson:"type"`
	Documentation string   `bson:"documentation"`
	Target        []string `bson:"target"`
	Chain         []string `bson:"chain"`
}

// This is an ugly hack to deal with embedded structures in the spec resource
type ConformanceRestResourceComponent struct {
	Type          string                                        `bson:"type"`
	Profile       Reference                                     `bson:"profile"`
	Interaction   []ResourceInteractionComponent                `bson:"interaction"`
	ReadHistory   bool                                          `bson:"readHistory"`
	UpdateCreate  bool                                          `bson:"updateCreate"`
	SearchInclude []string                                      `bson:"searchInclude"`
	SearchParam   []ConformanceRestResourceSearchParamComponent `bson:"searchParam"`
}

// This is an ugly hack to deal with embedded structures in the spec interaction
type SystemInteractionComponent struct {
	Code          string `bson:"code"`
	Documentation string `bson:"documentation"`
}

// This is an ugly hack to deal with embedded structures in the spec operation
type ConformanceRestOperationComponent struct {
	Name       string    `bson:"name"`
	Definition Reference `bson:"definition"`
}

// This is an ugly hack to deal with embedded structures in the spec rest
type ConformanceRestComponent struct {
	Mode            string                              `bson:"mode"`
	Documentation   string                              `bson:"documentation"`
	Security        ConformanceRestSecurityComponent    `bson:"security"`
	Resource        []ConformanceRestResourceComponent  `bson:"resource"`
	Interaction     []SystemInteractionComponent        `bson:"interaction"`
	Operation       []ConformanceRestOperationComponent `bson:"operation"`
	DocumentMailbox []string                            `bson:"documentMailbox"`
}

// This is an ugly hack to deal with embedded structures in the spec event
type ConformanceMessagingEventComponent struct {
	Code          Coding    `bson:"code"`
	Category      string    `bson:"category"`
	Mode          string    `bson:"mode"`
	Protocol      []Coding  `bson:"protocol"`
	Focus         string    `bson:"focus"`
	Request       Reference `bson:"request"`
	Response      Reference `bson:"response"`
	Documentation string    `bson:"documentation"`
}

// This is an ugly hack to deal with embedded structures in the spec messaging
type ConformanceMessagingComponent struct {
	Endpoint      string                               `bson:"endpoint"`
	ReliableCache float64                              `bson:"reliableCache"`
	Documentation string                               `bson:"documentation"`
	Event         []ConformanceMessagingEventComponent `bson:"event"`
}

// This is an ugly hack to deal with embedded structures in the spec document
type ConformanceDocumentComponent struct {
	Mode          string    `bson:"mode"`
	Documentation string    `bson:"documentation"`
	Profile       Reference `bson:"profile"`
}
