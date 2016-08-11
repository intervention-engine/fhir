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

type ModuleMetadata struct {
	Url             string                                   `bson:"url,omitempty" json:"url,omitempty"`
	Identifier      []Identifier                             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version         string                                   `bson:"version,omitempty" json:"version,omitempty"`
	Name            string                                   `bson:"name,omitempty" json:"name,omitempty"`
	Title           string                                   `bson:"title,omitempty" json:"title,omitempty"`
	Type            string                                   `bson:"type,omitempty" json:"type,omitempty"`
	Status          string                                   `bson:"status,omitempty" json:"status,omitempty"`
	Experimental    *bool                                    `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Description     string                                   `bson:"description,omitempty" json:"description,omitempty"`
	Purpose         string                                   `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage           string                                   `bson:"usage,omitempty" json:"usage,omitempty"`
	PublicationDate *FHIRDateTime                            `bson:"publicationDate,omitempty" json:"publicationDate,omitempty"`
	LastReviewDate  *FHIRDateTime                            `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod *Period                                  `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Coverage        []ModuleMetadataCoverageComponent        `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Topic           []CodeableConcept                        `bson:"topic,omitempty" json:"topic,omitempty"`
	Contributor     []ModuleMetadataContributorComponent     `bson:"contributor,omitempty" json:"contributor,omitempty"`
	Publisher       string                                   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact         []ModuleMetadataContactComponent         `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright       string                                   `bson:"copyright,omitempty" json:"copyright,omitempty"`
	RelatedResource []ModuleMetadataRelatedResourceComponent `bson:"relatedResource,omitempty" json:"relatedResource,omitempty"`
}

type ModuleMetadataCoverageComponent struct {
	BackboneElement `bson:",inline"`
	Focus           *Coding          `bson:"focus,omitempty" json:"focus,omitempty"`
	Value           *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
}

type ModuleMetadataContributorComponent struct {
	BackboneElement `bson:",inline"`
	Type            string                                      `bson:"type,omitempty" json:"type,omitempty"`
	Name            string                                      `bson:"name,omitempty" json:"name,omitempty"`
	Contact         []ModuleMetadataContributorContactComponent `bson:"contact,omitempty" json:"contact,omitempty"`
}

type ModuleMetadataContributorContactComponent struct {
	BackboneElement `bson:",inline"`
	Name            string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type ModuleMetadataContactComponent struct {
	BackboneElement `bson:",inline"`
	Name            string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type ModuleMetadataRelatedResourceComponent struct {
	BackboneElement `bson:",inline"`
	Type            string      `bson:"type,omitempty" json:"type,omitempty"`
	Document        *Attachment `bson:"document,omitempty" json:"document,omitempty"`
	Resource        *Reference  `bson:"resource,omitempty" json:"resource,omitempty"`
}
