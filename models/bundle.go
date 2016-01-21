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

type Bundle struct {
	Resource  `bson:",inline"`
	Type      string                 `bson:"type,omitempty" json:"type,omitempty"`
	Total     *uint32                `bson:"total,omitempty" json:"total,omitempty"`
	Link      []BundleLinkComponent  `bson:"link,omitempty" json:"link,omitempty"`
	Entry     []BundleEntryComponent `bson:"entry,omitempty" json:"entry,omitempty"`
	Signature *Signature             `bson:"signature,omitempty" json:"signature,omitempty"`
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
	Link     []BundleLinkComponent         `bson:"link,omitempty" json:"link,omitempty"`
	FullUrl  string                        `bson:"fullUrl,omitempty" json:"fullUrl,omitempty"`
	Resource interface{}                   `bson:"resource,omitempty" json:"resource,omitempty"`
	Search   *BundleEntrySearchComponent   `bson:"search,omitempty" json:"search,omitempty"`
	Request  *BundleEntryRequestComponent  `bson:"request,omitempty" json:"request,omitempty"`
	Response *BundleEntryResponseComponent `bson:"response,omitempty" json:"response,omitempty"`
}

// The "bundleEntryComponent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type bundleEntryComponent BundleEntryComponent

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *BundleEntryComponent) UnmarshalJSON(data []byte) (err error) {
	x2 := bundleEntryComponent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Resource != nil {
			x2.Resource = MapToResource(x2.Resource, true)
		}
		*x = BundleEntryComponent(x2)
	}
	return
}

type BundleEntrySearchComponent struct {
	Mode  string   `bson:"mode,omitempty" json:"mode,omitempty"`
	Score *float64 `bson:"score,omitempty" json:"score,omitempty"`
}

type BundleEntryRequestComponent struct {
	Method          string        `bson:"method,omitempty" json:"method,omitempty"`
	Url             string        `bson:"url,omitempty" json:"url,omitempty"`
	IfNoneMatch     string        `bson:"ifNoneMatch,omitempty" json:"ifNoneMatch,omitempty"`
	IfModifiedSince *FHIRDateTime `bson:"ifModifiedSince,omitempty" json:"ifModifiedSince,omitempty"`
	IfMatch         string        `bson:"ifMatch,omitempty" json:"ifMatch,omitempty"`
	IfNoneExist     string        `bson:"ifNoneExist,omitempty" json:"ifNoneExist,omitempty"`
}

type BundleEntryResponseComponent struct {
	Status       string        `bson:"status,omitempty" json:"status,omitempty"`
	Location     string        `bson:"location,omitempty" json:"location,omitempty"`
	Etag         string        `bson:"etag,omitempty" json:"etag,omitempty"`
	LastModified *FHIRDateTime `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
}

type BundlePlus struct {
	Bundle             `bson:",inline"`
	BundlePlusIncludes `bson:",inline"`
}

type BundlePlusIncludes struct {
	IncludedCompositionResources *[]Composition   `bson:"_includedCompositionResources,omitempty"`
	IncludedMessageResources     *[]MessageHeader `bson:"_includedMessageResources,omitempty"`
}

func (b *BundlePlusIncludes) GetIncludedCompositionResource() (composition *Composition, err error) {
	if b.IncludedCompositionResources == nil {
		err = errors.New("Included compositions not requested")
	} else if len(*b.IncludedCompositionResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 composition, but found %d", len(*b.IncludedCompositionResources))
	} else if len(*b.IncludedCompositionResources) == 1 {
		composition = &(*b.IncludedCompositionResources)[0]
	}
	return
}

func (b *BundlePlusIncludes) GetIncludedMessageResource() (messageHeader *MessageHeader, err error) {
	if b.IncludedMessageResources == nil {
		err = errors.New("Included messageheaders not requested")
	} else if len(*b.IncludedMessageResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 messageHeader, but found %d", len(*b.IncludedMessageResources))
	} else if len(*b.IncludedMessageResources) == 1 {
		messageHeader = &(*b.IncludedMessageResources)[0]
	}
	return
}

func (b *BundlePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.IncludedCompositionResources != nil {
		for _, r := range *b.IncludedCompositionResources {
			resourceMap[r.Id] = &r
		}
	}
	if b.IncludedMessageResources != nil {
		for _, r := range *b.IncludedMessageResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
