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

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Binary struct {
	Resource    `bson:",inline"`
	ContentType string `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Content     string `bson:"content,omitempty" json:"content,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Binary) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Binary"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Binary), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Binary) GetBSON() (interface{}, error) {
	x.ResourceType = "Binary"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "binary" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type binary Binary

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Binary) UnmarshalJSON(data []byte) (err error) {
	x2 := binary{}
	if err = json.Unmarshal(data, &x2); err == nil {
		*x = Binary(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Binary) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Binary"
	} else if x.ResourceType != "Binary" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Binary, instead received %s", x.ResourceType))
	}
	return nil
}

type BinaryPlus struct {
	Binary             `bson:",inline"`
	BinaryPlusIncludes `bson:",inline"`
}

type BinaryPlusIncludes struct {
}

func (b *BinaryPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}
