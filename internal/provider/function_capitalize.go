// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"

	tfplugin_string "terraform-provider-battery/internal/libs/string"
)

var _ function.Function = &CapitalizeFunction{}

type CapitalizeFunction struct{}

// please register this function name at internal/provider/provider.go on your TF providers
func NewCapitalizeFunction() function.Function {
	return &CapitalizeFunction{}
}

func (f *CapitalizeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "capitalize"
}

// You can modify the required parameters on the Definition func
func (f *CapitalizeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert string to capitalized",
		Description: "Convert string to capitalized",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "string",
				Description: "string to be converted",
			},
		},
		Return: function.StringReturn{},
	}
}

// You can modify the required parameters and operations on the Run func
func (f *CapitalizeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var str string

	// catch the params by using defined variables
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str))
	if resp.Error != nil {
		return
	}

	// main operation is written here. It's modifiable
	result, err := tfplugin_string.Capitalize(str)

	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Error executing function: "+err.Error()))
		return
	} else {
		// return the value of above operation by following this format
		resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &result))
	}
}
