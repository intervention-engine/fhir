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

type Bundle struct {
	Id        string                 `json:"-" bson:"_id"`
	Type      string                 `bson:"type,omitempty" json:"type,omitempty"`
	Base      string                 `bson:"base,omitempty" json:"base,omitempty"`
	Total     *uint32                `bson:"total,omitempty" json:"total,omitempty"`
	Link      []BundleLinkComponent  `bson:"link,omitempty" json:"link,omitempty"`
	Entry     []BundleEntryComponent `bson:"entry,omitempty" json:"entry,omitempty"`
	Signature string                 `bson:"signature,omitempty" json:"signature,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Bundle) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Bundle
	}{
		ResourceType: "Bundle",
		Bundle:       *resource,
	}
	return json.Marshal(x)
}

type BundleLinkComponent struct {
	Relation string `bson:"relation,omitempty" json:"relation,omitempty"`
	Url      string `bson:"url,omitempty" json:"url,omitempty"`
}

type BundleEntryComponent struct {
	Base                string                                   `bson:"base,omitempty" json:"base,omitempty"`
	Link                []BundleLinkComponent                    `bson:"link,omitempty" json:"link,omitempty"`
	Resource            interface{}                              `bson:"resource,omitempty" json:"resource,omitempty"`
	Search              *BundleEntrySearchComponent              `bson:"search,omitempty" json:"search,omitempty"`
	Transaction         *BundleEntryTransactionComponent         `bson:"transaction,omitempty" json:"transaction,omitempty"`
	TransactionResponse *BundleEntryTransactionResponseComponent `bson:"transactionResponse,omitempty" json:"transactionResponse,omitempty"`
}

// The "bundleEntryComponent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type bundleEntryComponent BundleEntryComponent

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *BundleEntryComponent) UnmarshalJSON(data []byte) (err error) {
	x2 := bundleEntryComponent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		x2.Resource = MapToResource(x2.Resource)
		*x = BundleEntryComponent(x2)
	}
	return
}

type BundleEntrySearchComponent struct {
	Mode  string   `bson:"mode,omitempty" json:"mode,omitempty"`
	Score *float64 `bson:"score,omitempty" json:"score,omitempty"`
}

type BundleEntryTransactionComponent struct {
	Method          string        `bson:"method,omitempty" json:"method,omitempty"`
	Url             string        `bson:"url,omitempty" json:"url,omitempty"`
	IfNoneMatch     string        `bson:"ifNoneMatch,omitempty" json:"ifNoneMatch,omitempty"`
	IfMatch         string        `bson:"ifMatch,omitempty" json:"ifMatch,omitempty"`
	IfModifiedSince *FHIRDateTime `bson:"ifModifiedSince,omitempty" json:"ifModifiedSince,omitempty"`
	IfNoneExist     string        `bson:"ifNoneExist,omitempty" json:"ifNoneExist,omitempty"`
}

type BundleEntryTransactionResponseComponent struct {
	Status       string        `bson:"status,omitempty" json:"status,omitempty"`
	Location     string        `bson:"location,omitempty" json:"location,omitempty"`
	Etag         string        `bson:"etag,omitempty" json:"etag,omitempty"`
	LastModified *FHIRDateTime `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
}
