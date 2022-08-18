// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import low "github.com/pb33f/libopenapi/datamodel/low/3.0"

type Parameter struct {
	Name            string
	In              string
	Description     string
	Required        bool
	Deprecated      bool
	AllowEmptyValue bool
	Style           string
	Explode         bool
	AllowReserved   bool
	Schema          *Schema
	Example         any
	Examples        map[string]*Example
	Content         map[string]*MediaType
	Extensions      map[string]any
	low             *low.Parameter
}

func (p *Parameter) GoLow() *low.Parameter {
	return p.low
}
