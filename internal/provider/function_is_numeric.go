// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"

	tfplugin_string "terraform-provider-battery/internal/libs/string"
)

var _ function.Function = &IsNumericFunction{}

type IsNumericFunction struct{}

// please register this function name at internal/provider/provider.go on your TF providers
func NewIsNumericFunction() function.Function {
	return &IsNumericFunction{}
}

func (f *IsNumericFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "is_numeric"
}

// You can modify the required parameters on the Definition func
func (f *IsNumericFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Check if given string is containing numeric only",
		Description: "Detect numeric 0-9 on given string",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "string",
				Description: "string to be checked",
			},
		},
		Return: function.BoolReturn{},
	}
}

// You can modify the required parameters and operations on the Run func
func (f *IsNumericFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var str string

	// catch the params by using defined variables
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str))
	if resp.Error != nil {
		return
	}

	// main operation is written here. It's modifiable
	result := tfplugin_string.IsNumeric(str)

	// return the value of above operation by following this format
	resp.Error = resp.Result.Set(ctx, &result)
}
