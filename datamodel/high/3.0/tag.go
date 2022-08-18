// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import low "github.com/pb33f/libopenapi/datamodel/low/3.0"

type Tag struct {
	Name         string
	Description  string
	ExternalDocs *ExternalDoc
	Extensions   map[string]any
	low          *low.Tag
}

func (t *Tag) GoLow() *low.Tag {
	return t.low
}