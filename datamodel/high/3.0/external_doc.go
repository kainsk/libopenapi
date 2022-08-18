// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import low "github.com/pb33f/libopenapi/datamodel/low/3.0"

type ExternalDoc struct {
	Description string
	URL         string
	Extensions  map[string]any
	low         *low.ExternalDoc
}

func (e *ExternalDoc) GoLow() *low.ExternalDoc {
	return e.low
}