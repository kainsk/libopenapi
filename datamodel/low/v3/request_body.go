// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"gopkg.in/yaml.v3"
)

type RequestBody struct {
	Description low.NodeReference[string]
	Content     low.NodeReference[map[low.KeyReference[string]]low.ValueReference[*MediaType]]
	Required    low.NodeReference[bool]
	Extensions  map[low.KeyReference[string]]low.ValueReference[any]
}

func (rb *RequestBody) FindExtension(ext string) *low.ValueReference[any] {
	return low.FindItemInMap[any](ext, rb.Extensions)
}

func (rb *RequestBody) FindContent(cType string) *low.ValueReference[*MediaType] {
	return low.FindItemInMap[*MediaType](cType, rb.Content.Value)
}

func (rb *RequestBody) Build(root *yaml.Node, idx *index.SpecIndex) error {
	rb.Extensions = low.ExtractExtensions(root)

	// handle content, if set.
	con, cL, cN, cErr := low.ExtractMap[*MediaType](ContentLabel, root, idx)
	if cErr != nil {
		return cErr
	}
	if con != nil {
		rb.Content = low.NodeReference[map[low.KeyReference[string]]low.ValueReference[*MediaType]]{
			Value:     con,
			KeyNode:   cL,
			ValueNode: cN,
		}
	}
	return nil
}