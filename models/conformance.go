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

type Conformance struct {
	Id             string                             `json:"-" bson:"_id"`
	Identifier     string                             `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Version        string                             `bson:"version,omitempty", json:"version,omitempty"`
	Name           string                             `bson:"name,omitempty", json:"name,omitempty"`
	Publisher      string                             `bson:"publisher,omitempty", json:"publisher,omitempty"`
	Telecom        []ContactPoint                     `bson:"telecom,omitempty", json:"telecom,omitempty"`
	Description    string                             `bson:"description,omitempty", json:"description,omitempty"`
	Status         string                             `bson:"status,omitempty", json:"status,omitempty"`
	Experimental   bool                               `bson:"experimental,omitempty", json:"experimental,omitempty"`
	Date           FHIRDateTime                       `bson:"date,omitempty", json:"date,omitempty"`
	Software       ConformanceSoftwareComponent       `bson:"software,omitempty", json:"software,omitempty"`
	Implementation ConformanceImplementationComponent `bson:"implementation,omitempty", json:"implementation,omitempty"`
	FhirVersion    string                             `bson:"fhirVersion,omitempty", json:"fhirVersion,omitempty"`
	AcceptUnknown  bool                               `bson:"acceptUnknown,omitempty", json:"acceptUnknown,omitempty"`
	Format         []string                           `bson:"format,omitempty", json:"format,omitempty"`
	Profile        []Reference                        `bson:"profile,omitempty", json:"profile,omitempty"`
	Rest           []ConformanceRestComponent         `bson:"rest,omitempty", json:"rest,omitempty"`
	Messaging      []ConformanceMessagingComponent    `bson:"messaging,omitempty", json:"messaging,omitempty"`
	Document       []ConformanceDocumentComponent     `bson:"document,omitempty", json:"document,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec software
type ConformanceSoftwareComponent struct {
	Name        string       `bson:"name,omitempty", json:"name,omitempty"`
	Version     string       `bson:"version,omitempty", json:"version,omitempty"`
	ReleaseDate FHIRDateTime `bson:"releaseDate,omitempty", json:"releaseDate,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec implementation
type ConformanceImplementationComponent struct {
	Description string `bson:"description,omitempty", json:"description,omitempty"`
	Url         string `bson:"url,omitempty", json:"url,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec certificate
type ConformanceRestSecurityCertificateComponent struct {
	Type string `bson:"type,omitempty", json:"type,omitempty"`
	Blob string `bson:"blob,omitempty", json:"blob,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec security
type ConformanceRestSecurityComponent struct {
	Cors        bool                                          `bson:"cors,omitempty", json:"cors,omitempty"`
	Service     []CodeableConcept                             `bson:"service,omitempty", json:"service,omitempty"`
	Description string                                        `bson:"description,omitempty", json:"description,omitempty"`
	Certificate []ConformanceRestSecurityCertificateComponent `bson:"certificate,omitempty", json:"certificate,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec interaction
type ResourceInteractionComponent struct {
	Code          string `bson:"code,omitempty", json:"code,omitempty"`
	Documentation string `bson:"documentation,omitempty", json:"documentation,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec searchParam
type ConformanceRestResourceSearchParamComponent struct {
	Name          string   `bson:"name,omitempty", json:"name,omitempty"`
	Definition    string   `bson:"definition,omitempty", json:"definition,omitempty"`
	Type          string   `bson:"type,omitempty", json:"type,omitempty"`
	Documentation string   `bson:"documentation,omitempty", json:"documentation,omitempty"`
	Target        []string `bson:"target,omitempty", json:"target,omitempty"`
	Chain         []string `bson:"chain,omitempty", json:"chain,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec resource
type ConformanceRestResourceComponent struct {
	Type          string                                        `bson:"type,omitempty", json:"type,omitempty"`
	Profile       Reference                                     `bson:"profile,omitempty", json:"profile,omitempty"`
	Interaction   []ResourceInteractionComponent                `bson:"interaction,omitempty", json:"interaction,omitempty"`
	ReadHistory   bool                                          `bson:"readHistory,omitempty", json:"readHistory,omitempty"`
	UpdateCreate  bool                                          `bson:"updateCreate,omitempty", json:"updateCreate,omitempty"`
	SearchInclude []string                                      `bson:"searchInclude,omitempty", json:"searchInclude,omitempty"`
	SearchParam   []ConformanceRestResourceSearchParamComponent `bson:"searchParam,omitempty", json:"searchParam,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec interaction
type SystemInteractionComponent struct {
	Code          string `bson:"code,omitempty", json:"code,omitempty"`
	Documentation string `bson:"documentation,omitempty", json:"documentation,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec operation
type ConformanceRestOperationComponent struct {
	Name       string    `bson:"name,omitempty", json:"name,omitempty"`
	Definition Reference `bson:"definition,omitempty", json:"definition,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec rest
type ConformanceRestComponent struct {
	Mode            string                              `bson:"mode,omitempty", json:"mode,omitempty"`
	Documentation   string                              `bson:"documentation,omitempty", json:"documentation,omitempty"`
	Security        ConformanceRestSecurityComponent    `bson:"security,omitempty", json:"security,omitempty"`
	Resource        []ConformanceRestResourceComponent  `bson:"resource,omitempty", json:"resource,omitempty"`
	Interaction     []SystemInteractionComponent        `bson:"interaction,omitempty", json:"interaction,omitempty"`
	Operation       []ConformanceRestOperationComponent `bson:"operation,omitempty", json:"operation,omitempty"`
	DocumentMailbox []string                            `bson:"documentMailbox,omitempty", json:"documentMailbox,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec event
type ConformanceMessagingEventComponent struct {
	Code          Coding    `bson:"code,omitempty", json:"code,omitempty"`
	Category      string    `bson:"category,omitempty", json:"category,omitempty"`
	Mode          string    `bson:"mode,omitempty", json:"mode,omitempty"`
	Protocol      []Coding  `bson:"protocol,omitempty", json:"protocol,omitempty"`
	Focus         string    `bson:"focus,omitempty", json:"focus,omitempty"`
	Request       Reference `bson:"request,omitempty", json:"request,omitempty"`
	Response      Reference `bson:"response,omitempty", json:"response,omitempty"`
	Documentation string    `bson:"documentation,omitempty", json:"documentation,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec messaging
type ConformanceMessagingComponent struct {
	Endpoint      string                               `bson:"endpoint,omitempty", json:"endpoint,omitempty"`
	ReliableCache float64                              `bson:"reliableCache,omitempty", json:"reliableCache,omitempty"`
	Documentation string                               `bson:"documentation,omitempty", json:"documentation,omitempty"`
	Event         []ConformanceMessagingEventComponent `bson:"event,omitempty", json:"event,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec document
type ConformanceDocumentComponent struct {
	Mode          string    `bson:"mode,omitempty", json:"mode,omitempty"`
	Documentation string    `bson:"documentation,omitempty", json:"documentation,omitempty"`
	Profile       Reference `bson:"profile,omitempty", json:"profile,omitempty"`
}
