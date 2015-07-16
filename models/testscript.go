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

import "encoding/json"

type TestScript struct {
	Id          string                       `json:"id" bson:"_id"`
	Name        string                       `bson:"name,omitempty" json:"name,omitempty"`
	Description string                       `bson:"description,omitempty" json:"description,omitempty"`
	Multiserver *bool                        `bson:"multiserver,omitempty" json:"multiserver,omitempty"`
	Fixture     []TestScriptFixtureComponent `bson:"fixture,omitempty" json:"fixture,omitempty"`
	Setup       *TestScriptSetupComponent    `bson:"setup,omitempty" json:"setup,omitempty"`
	Test        []TestScriptTestComponent    `bson:"test,omitempty" json:"test,omitempty"`
	Teardown    *TestScriptTeardownComponent `bson:"teardown,omitempty" json:"teardown,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *TestScript) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		TestScript
	}{
		ResourceType: "TestScript",
		TestScript:   *resource,
	}
	return json.Marshal(x)
}

type TestScriptFixtureComponent struct {
	Uri        string      `bson:"uri,omitempty" json:"uri,omitempty"`
	Resource   interface{} `bson:"resource,omitempty" json:"resource,omitempty"`
	Autocreate *bool       `bson:"autocreate,omitempty" json:"autocreate,omitempty"`
	Autodelete *bool       `bson:"autodelete,omitempty" json:"autodelete,omitempty"`
}

// The "testScriptFixtureComponent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type testScriptFixtureComponent TestScriptFixtureComponent

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *TestScriptFixtureComponent) UnmarshalJSON(data []byte) (err error) {
	x2 := testScriptFixtureComponent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		x2.Resource = MapToResource(x2.Resource)
		*x = TestScriptFixtureComponent(x2)
	}
	return
}

type TestScriptSetupComponent struct {
	Operation []TestScriptSetupOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
}

type TestScriptSetupOperationComponent struct {
	Type        string   `bson:"type,omitempty" json:"type,omitempty"`
	Source      string   `bson:"source,omitempty" json:"source,omitempty"`
	Target      string   `bson:"target,omitempty" json:"target,omitempty"`
	Destination *int32   `bson:"destination,omitempty" json:"destination,omitempty"`
	Parameter   []string `bson:"parameter,omitempty" json:"parameter,omitempty"`
	ResponseId  string   `bson:"responseId,omitempty" json:"responseId,omitempty"`
	ContentType string   `bson:"contentType,omitempty" json:"contentType,omitempty"`
}

type TestScriptTestComponent struct {
	Name        string                             `bson:"name,omitempty" json:"name,omitempty"`
	Description string                             `bson:"description,omitempty" json:"description,omitempty"`
	Metadata    *TestScriptTestMetadataComponent   `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Operation   []TestScriptTestOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
}

type TestScriptTestMetadataComponent struct {
	Link      []TestScriptTestMetadataLinkComponent      `bson:"link,omitempty" json:"link,omitempty"`
	Requires  []TestScriptTestMetadataRequiresComponent  `bson:"requires,omitempty" json:"requires,omitempty"`
	Validates []TestScriptTestMetadataValidatesComponent `bson:"validates,omitempty" json:"validates,omitempty"`
}

type TestScriptTestMetadataLinkComponent struct {
	Url         string `bson:"url,omitempty" json:"url,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}

type TestScriptTestMetadataRequiresComponent struct {
	Type        string `bson:"type,omitempty" json:"type,omitempty"`
	Operations  string `bson:"operations,omitempty" json:"operations,omitempty"`
	Destination *int32 `bson:"destination,omitempty" json:"destination,omitempty"`
}

type TestScriptTestMetadataValidatesComponent struct {
	Type        string `bson:"type,omitempty" json:"type,omitempty"`
	Operations  string `bson:"operations,omitempty" json:"operations,omitempty"`
	Destination *int32 `bson:"destination,omitempty" json:"destination,omitempty"`
}

type TestScriptTestOperationComponent struct {
	Type        string   `bson:"type,omitempty" json:"type,omitempty"`
	Source      string   `bson:"source,omitempty" json:"source,omitempty"`
	Target      string   `bson:"target,omitempty" json:"target,omitempty"`
	Destination *int32   `bson:"destination,omitempty" json:"destination,omitempty"`
	Parameter   []string `bson:"parameter,omitempty" json:"parameter,omitempty"`
	ResponseId  string   `bson:"responseId,omitempty" json:"responseId,omitempty"`
	ContentType string   `bson:"contentType,omitempty" json:"contentType,omitempty"`
}

type TestScriptTeardownComponent struct {
	Operation []TestScriptTeardownOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
}

type TestScriptTeardownOperationComponent struct {
	Type        string   `bson:"type,omitempty" json:"type,omitempty"`
	Source      string   `bson:"source,omitempty" json:"source,omitempty"`
	Target      string   `bson:"target,omitempty" json:"target,omitempty"`
	Destination *int32   `bson:"destination,omitempty" json:"destination,omitempty"`
	Parameter   []string `bson:"parameter,omitempty" json:"parameter,omitempty"`
	ResponseId  string   `bson:"responseId,omitempty" json:"responseId,omitempty"`
	ContentType string   `bson:"contentType,omitempty" json:"contentType,omitempty"`
}
