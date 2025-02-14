// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	tfplugin_string "terraform-provider-battery/internal/libs/string"
)

var _ function.Function = &SnakeCaseFunction{}

type SnakeCaseFunction struct{}

// please register this function name at internal/provider/provider.go on your TF providers
func NewSnakeCaseFunction() function.Function {
	return &SnakeCaseFunction{}
}

func (f *SnakeCaseFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "snake_case"
}

// You can modify the required parameters on the Definition func
func (f *SnakeCaseFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert given string to Snake Case",
		Description: "Convert given string to Snake Case by using Go Standard Library",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "string",
				Description: "string to be converted to Snake Case",
			},
		},
		Return: function.StringReturn{},
	}
}

// You can modify the required parameters and operations on the Run func
func (f *SnakeCaseFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var str string

	// catch the params by using defined variables
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str))
	if resp.Error != nil {
		return
	}

	// main operation is written here. It's modifiable
	result := tfplugin_string.SnakeCase(str)

	// return the value of above operation by following this format
	resp.Error = resp.Result.Set(ctx, &result)
}
