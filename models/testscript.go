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

type TestScript struct {
	Id          string                                 `json:"-" bson:"_id"`
	Name        string                                 `bson:"name,omitempty" json:"name,omitempty"`
	Description string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Multiserver *bool                                  `bson:"multiserver,omitempty" json:"multiserver,omitempty"`
	Fixture     []TestScriptTestScriptFixtureComponent `bson:"fixture,omitempty" json:"fixture,omitempty"`
	Setup       *TestScriptTestScriptSetupComponent    `bson:"setup,omitempty" json:"setup,omitempty"`
	Test        []TestScriptTestScriptTestComponent    `bson:"test,omitempty" json:"test,omitempty"`
	Teardown    *TestScriptTestScriptTeardownComponent `bson:"teardown,omitempty" json:"teardown,omitempty"`
}
type TestScriptTestScriptFixtureComponent struct {
	Uri        string     `bson:"uri,omitempty" json:"uri,omitempty"`
	Resource   *Reference `bson:"resource,omitempty" json:"resource,omitempty"`
	Autocreate *bool      `bson:"autocreate,omitempty" json:"autocreate,omitempty"`
	Autodelete *bool      `bson:"autodelete,omitempty" json:"autodelete,omitempty"`
}
type TestScriptTestScriptSetupComponent struct {
	Operation []TestScriptTestScriptSetupOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
}
type TestScriptTestScriptSetupOperationComponent struct {
	Type        string   `bson:"type,omitempty" json:"type,omitempty"`
	Source      string   `bson:"source,omitempty" json:"source,omitempty"`
	Target      string   `bson:"target,omitempty" json:"target,omitempty"`
	Destination *int32   `bson:"destination,omitempty" json:"destination,omitempty"`
	Parameter   []string `bson:"parameter,omitempty" json:"parameter,omitempty"`
	ResponseId  string   `bson:"responseId,omitempty" json:"responseId,omitempty"`
	ContentType string   `bson:"contentType,omitempty" json:"contentType,omitempty"`
}
type TestScriptTestScriptTestComponent struct {
	Name        string                                       `bson:"name,omitempty" json:"name,omitempty"`
	Description string                                       `bson:"description,omitempty" json:"description,omitempty"`
	Metadata    *TestScriptTestScriptTestMetadataComponent   `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Operation   []TestScriptTestScriptTestOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
}
type TestScriptTestScriptTestMetadataComponent struct {
	Link      []TestScriptTestScriptTestMetadataLinkComponent      `bson:"link,omitempty" json:"link,omitempty"`
	Requires  []TestScriptTestScriptTestMetadataRequiresComponent  `bson:"requires,omitempty" json:"requires,omitempty"`
	Validates []TestScriptTestScriptTestMetadataValidatesComponent `bson:"validates,omitempty" json:"validates,omitempty"`
}
type TestScriptTestScriptTestMetadataLinkComponent struct {
	Url         string `bson:"url,omitempty" json:"url,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}
type TestScriptTestScriptTestMetadataRequiresComponent struct {
	Type        string `bson:"type,omitempty" json:"type,omitempty"`
	Operations  string `bson:"operations,omitempty" json:"operations,omitempty"`
	Destination *int32 `bson:"destination,omitempty" json:"destination,omitempty"`
}
type TestScriptTestScriptTestMetadataValidatesComponent struct {
	Type        string `bson:"type,omitempty" json:"type,omitempty"`
	Operations  string `bson:"operations,omitempty" json:"operations,omitempty"`
	Destination *int32 `bson:"destination,omitempty" json:"destination,omitempty"`
}
type TestScriptTestScriptTestOperationComponent struct {
	Type        string   `bson:"type,omitempty" json:"type,omitempty"`
	Source      string   `bson:"source,omitempty" json:"source,omitempty"`
	Target      string   `bson:"target,omitempty" json:"target,omitempty"`
	Destination *int32   `bson:"destination,omitempty" json:"destination,omitempty"`
	Parameter   []string `bson:"parameter,omitempty" json:"parameter,omitempty"`
	ResponseId  string   `bson:"responseId,omitempty" json:"responseId,omitempty"`
	ContentType string   `bson:"contentType,omitempty" json:"contentType,omitempty"`
}
type TestScriptTestScriptTeardownComponent struct {
	Operation []TestScriptTestScriptTeardownOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
}
type TestScriptTestScriptTeardownOperationComponent struct {
	Type        string   `bson:"type,omitempty" json:"type,omitempty"`
	Source      string   `bson:"source,omitempty" json:"source,omitempty"`
	Target      string   `bson:"target,omitempty" json:"target,omitempty"`
	Destination *int32   `bson:"destination,omitempty" json:"destination,omitempty"`
	Parameter   []string `bson:"parameter,omitempty" json:"parameter,omitempty"`
	ResponseId  string   `bson:"responseId,omitempty" json:"responseId,omitempty"`
	ContentType string   `bson:"contentType,omitempty" json:"contentType,omitempty"`
}

type TestScriptBundle struct {
	Type         string                  `json:"resourceType,omitempty"`
	Title        string                  `json:"title,omitempty"`
	Id           string                  `json:"id,omitempty"`
	Updated      time.Time               `json:"updated,omitempty"`
	TotalResults int                     `json:"totalResults,omitempty"`
	Entry        []TestScriptBundleEntry `json:"entry,omitempty"`
	Category     TestScriptCategory      `json:"category,omitempty"`
}

type TestScriptBundleEntry struct {
	Title    string             `json:"title,omitempty"`
	Id       string             `json:"id,omitempty"`
	Content  TestScript         `json:"content,omitempty"`
	Category TestScriptCategory `json:"category,omitempty"`
}

type TestScriptCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
