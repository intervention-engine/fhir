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

type ActionDefinition struct {
	ActionIdentifier   *Identifier                              `bson:"actionIdentifier,omitempty" json:"actionIdentifier,omitempty"`
	Label              string                                   `bson:"label,omitempty" json:"label,omitempty"`
	Title              string                                   `bson:"title,omitempty" json:"title,omitempty"`
	Description        string                                   `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent     string                                   `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	Concept            []CodeableConcept                        `bson:"concept,omitempty" json:"concept,omitempty"`
	SupportingEvidence []Attachment                             `bson:"supportingEvidence,omitempty" json:"supportingEvidence,omitempty"`
	Documentation      []Attachment                             `bson:"documentation,omitempty" json:"documentation,omitempty"`
	RelatedAction      *ActionDefinitionRelatedActionComponent  `bson:"relatedAction,omitempty" json:"relatedAction,omitempty"`
	ParticipantType    []string                                 `bson:"participantType,omitempty" json:"participantType,omitempty"`
	Type               string                                   `bson:"type,omitempty" json:"type,omitempty"`
	Behavior           []ActionDefinitionBehaviorComponent      `bson:"behavior,omitempty" json:"behavior,omitempty"`
	Resource           *Reference                               `bson:"resource,omitempty" json:"resource,omitempty"`
	Customization      []ActionDefinitionCustomizationComponent `bson:"customization,omitempty" json:"customization,omitempty"`
	Action             []ActionDefinition                       `bson:"action,omitempty" json:"action,omitempty"`
}

type ActionDefinitionRelatedActionComponent struct {
	BackboneElement  `bson:",inline"`
	ActionIdentifier *Identifier `bson:"actionIdentifier,omitempty" json:"actionIdentifier,omitempty"`
	Relationship     string      `bson:"relationship,omitempty" json:"relationship,omitempty"`
	OffsetDuration   *Quantity   `bson:"offsetDuration,omitempty" json:"offsetDuration,omitempty"`
	OffsetRange      *Range      `bson:"offsetRange,omitempty" json:"offsetRange,omitempty"`
	Anchor           string      `bson:"anchor,omitempty" json:"anchor,omitempty"`
}

type ActionDefinitionBehaviorComponent struct {
	BackboneElement `bson:",inline"`
	Type            *Coding `bson:"type,omitempty" json:"type,omitempty"`
	Value           *Coding `bson:"value,omitempty" json:"value,omitempty"`
}

type ActionDefinitionCustomizationComponent struct {
	BackboneElement `bson:",inline"`
	Path            string `bson:"path,omitempty" json:"path,omitempty"`
	Expression      string `bson:"expression,omitempty" json:"expression,omitempty"`
}
