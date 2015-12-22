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
	DomainResource `bson:",inline"`
	Url            string                        `bson:"url,omitempty" json:"url,omitempty"`
	Version        string                        `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                        `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                        `bson:"status,omitempty" json:"status,omitempty"`
	Identifier     *Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Experimental   *bool                         `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                        `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []TestScriptContactComponent  `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                 `bson:"date,omitempty" json:"date,omitempty"`
	Description    string                        `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []CodeableConcept             `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Requirements   string                        `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright      string                        `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Metadata       *TestScriptMetadataComponent  `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Multiserver    *bool                         `bson:"multiserver,omitempty" json:"multiserver,omitempty"`
	Fixture        []TestScriptFixtureComponent  `bson:"fixture,omitempty" json:"fixture,omitempty"`
	Profile        []Reference                   `bson:"profile,omitempty" json:"profile,omitempty"`
	Variable       []TestScriptVariableComponent `bson:"variable,omitempty" json:"variable,omitempty"`
	Setup          *TestScriptSetupComponent     `bson:"setup,omitempty" json:"setup,omitempty"`
	Test           []TestScriptTestComponent     `bson:"test,omitempty" json:"test,omitempty"`
	Teardown       *TestScriptTeardownComponent  `bson:"teardown,omitempty" json:"teardown,omitempty"`
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

// The "testScript" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type testScript TestScript

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *TestScript) UnmarshalJSON(data []byte) (err error) {
	x2 := testScript{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = TestScript(x2)
	}
	return
}

type TestScriptContactComponent struct {
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type TestScriptMetadataComponent struct {
	Link       []TestScriptMetadataLinkComponent       `bson:"link,omitempty" json:"link,omitempty"`
	Capability []TestScriptMetadataCapabilityComponent `bson:"capability,omitempty" json:"capability,omitempty"`
}

type TestScriptMetadataLinkComponent struct {
	Url         string `bson:"url,omitempty" json:"url,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}

type TestScriptMetadataCapabilityComponent struct {
	Required    *bool      `bson:"required,omitempty" json:"required,omitempty"`
	Validated   *bool      `bson:"validated,omitempty" json:"validated,omitempty"`
	Description string     `bson:"description,omitempty" json:"description,omitempty"`
	Destination *int32     `bson:"destination,omitempty" json:"destination,omitempty"`
	Link        []string   `bson:"link,omitempty" json:"link,omitempty"`
	Conformance *Reference `bson:"conformance,omitempty" json:"conformance,omitempty"`
}

type TestScriptFixtureComponent struct {
	Autocreate *bool      `bson:"autocreate,omitempty" json:"autocreate,omitempty"`
	Autodelete *bool      `bson:"autodelete,omitempty" json:"autodelete,omitempty"`
	Resource   *Reference `bson:"resource,omitempty" json:"resource,omitempty"`
}

type TestScriptVariableComponent struct {
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	HeaderField string `bson:"headerField,omitempty" json:"headerField,omitempty"`
	Path        string `bson:"path,omitempty" json:"path,omitempty"`
	SourceId    string `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
}

type TestScriptSetupComponent struct {
	Metadata *TestScriptMetadataComponent     `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Action   []TestScriptSetupActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestScriptSetupActionComponent struct {
	Operation *TestScriptSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert    *TestScriptSetupActionAssertComponent    `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestScriptSetupActionOperationComponent struct {
	Type             *Coding                                                `bson:"type,omitempty" json:"type,omitempty"`
	Resource         string                                                 `bson:"resource,omitempty" json:"resource,omitempty"`
	Label            string                                                 `bson:"label,omitempty" json:"label,omitempty"`
	Description      string                                                 `bson:"description,omitempty" json:"description,omitempty"`
	Accept           string                                                 `bson:"accept,omitempty" json:"accept,omitempty"`
	ContentType      string                                                 `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Destination      *int32                                                 `bson:"destination,omitempty" json:"destination,omitempty"`
	EncodeRequestUrl *bool                                                  `bson:"encodeRequestUrl,omitempty" json:"encodeRequestUrl,omitempty"`
	Params           string                                                 `bson:"params,omitempty" json:"params,omitempty"`
	RequestHeader    []TestScriptSetupActionOperationRequestHeaderComponent `bson:"requestHeader,omitempty" json:"requestHeader,omitempty"`
	ResponseId       string                                                 `bson:"responseId,omitempty" json:"responseId,omitempty"`
	SourceId         string                                                 `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	TargetId         string                                                 `bson:"targetId,omitempty" json:"targetId,omitempty"`
	Url              string                                                 `bson:"url,omitempty" json:"url,omitempty"`
}

type TestScriptSetupActionOperationRequestHeaderComponent struct {
	Field string `bson:"field,omitempty" json:"field,omitempty"`
	Value string `bson:"value,omitempty" json:"value,omitempty"`
}

type TestScriptSetupActionAssertComponent struct {
	Label               string `bson:"label,omitempty" json:"label,omitempty"`
	Description         string `bson:"description,omitempty" json:"description,omitempty"`
	Direction           string `bson:"direction,omitempty" json:"direction,omitempty"`
	CompareToSourceId   string `bson:"compareToSourceId,omitempty" json:"compareToSourceId,omitempty"`
	CompareToSourcePath string `bson:"compareToSourcePath,omitempty" json:"compareToSourcePath,omitempty"`
	ContentType         string `bson:"contentType,omitempty" json:"contentType,omitempty"`
	HeaderField         string `bson:"headerField,omitempty" json:"headerField,omitempty"`
	MinimumId           string `bson:"minimumId,omitempty" json:"minimumId,omitempty"`
	NavigationLinks     *bool  `bson:"navigationLinks,omitempty" json:"navigationLinks,omitempty"`
	Operator            string `bson:"operator,omitempty" json:"operator,omitempty"`
	Path                string `bson:"path,omitempty" json:"path,omitempty"`
	Resource            string `bson:"resource,omitempty" json:"resource,omitempty"`
	Response            string `bson:"response,omitempty" json:"response,omitempty"`
	ResponseCode        string `bson:"responseCode,omitempty" json:"responseCode,omitempty"`
	SourceId            string `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	ValidateProfileId   string `bson:"validateProfileId,omitempty" json:"validateProfileId,omitempty"`
	Value               string `bson:"value,omitempty" json:"value,omitempty"`
	WarningOnly         *bool  `bson:"warningOnly,omitempty" json:"warningOnly,omitempty"`
}

type TestScriptTestComponent struct {
	Name        string                          `bson:"name,omitempty" json:"name,omitempty"`
	Description string                          `bson:"description,omitempty" json:"description,omitempty"`
	Metadata    *TestScriptMetadataComponent    `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Action      []TestScriptTestActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestScriptTestActionComponent struct {
	Operation *TestScriptSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert    *TestScriptSetupActionAssertComponent    `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestScriptTeardownComponent struct {
	Action []TestScriptTeardownActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestScriptTeardownActionComponent struct {
	Operation *TestScriptSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
}
