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

type ConceptMap struct {
	Id              string                       `json:"-" bson:"_id"`
	Identifier      string                       `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Version         string                       `bson:"version,omitempty", json:"version,omitempty"`
	Name            string                       `bson:"name,omitempty", json:"name,omitempty"`
	Publisher       string                       `bson:"publisher,omitempty", json:"publisher,omitempty"`
	Telecom         []ContactPoint               `bson:"telecom,omitempty", json:"telecom,omitempty"`
	Description     string                       `bson:"description,omitempty", json:"description,omitempty"`
	Copyright       string                       `bson:"copyright,omitempty", json:"copyright,omitempty"`
	Status          string                       `bson:"status,omitempty", json:"status,omitempty"`
	Experimental    bool                         `bson:"experimental,omitempty", json:"experimental,omitempty"`
	Date            FHIRDateTime                 `bson:"date,omitempty", json:"date,omitempty"`
	SourceUri       string                       `bson:"sourceUri,omitempty", json:"sourceUri,omitempty"`
	SourceReference Reference                    `bson:"sourceReference,omitempty", json:"sourceReference,omitempty"`
	TargetUri       string                       `bson:"targetUri,omitempty", json:"targetUri,omitempty"`
	TargetReference Reference                    `bson:"targetReference,omitempty", json:"targetReference,omitempty"`
	Element         []ConceptMapElementComponent `bson:"element,omitempty", json:"element,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec dependsOn
type OtherElementComponent struct {
	Element    string `bson:"element,omitempty", json:"element,omitempty"`
	CodeSystem string `bson:"codeSystem,omitempty", json:"codeSystem,omitempty"`
	Code       string `bson:"code,omitempty", json:"code,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec map
type ConceptMapElementMapComponent struct {
	CodeSystem  string                  `bson:"codeSystem,omitempty", json:"codeSystem,omitempty"`
	Code        string                  `bson:"code,omitempty", json:"code,omitempty"`
	Equivalence string                  `bson:"equivalence,omitempty", json:"equivalence,omitempty"`
	Comments    string                  `bson:"comments,omitempty", json:"comments,omitempty"`
	Product     []OtherElementComponent `bson:"product,omitempty", json:"product,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec element
type ConceptMapElementComponent struct {
	CodeSystem string                          `bson:"codeSystem,omitempty", json:"codeSystem,omitempty"`
	Code       string                          `bson:"code,omitempty", json:"code,omitempty"`
	DependsOn  []OtherElementComponent         `bson:"dependsOn,omitempty", json:"dependsOn,omitempty"`
	Map        []ConceptMapElementMapComponent `bson:"map,omitempty", json:"map,omitempty"`
}
