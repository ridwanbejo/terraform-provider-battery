// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"

	tfplugin_string "terraform-provider-battery/internal/libs/string"
)

var _ function.Function = &StrInsertFunction{}

type StrInsertFunction struct{}

// please register this function name at internal/provider/provider.go on your TF providers
func NewStrInsertFunction() function.Function {
	return &StrInsertFunction{}
}

func (f *StrInsertFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "str_insert"
}

// You can modify the required parameters on the Definition func
func (f *StrInsertFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Add substring on desired index to given string",
		Description: "Add substring on desired index to given string",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "string",
				Description: "string to be inserted",
			},
			function.StringParameter{
				Name:        "substring",
				Description: "substring to be inserted to given string",
			},
			function.Int32Parameter{
				Name:        "index",
				Description: "index where substring will be inserted",
			},
		},
		Return: function.StringReturn{},
	}
}

// You can modify the required parameters and operations on the Run func
func (f *StrInsertFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var str string
	var substr string
	var index int

	// catch the params by using defined variables
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str, &substr, &index))
	if resp.Error != nil {
		return
	}

	// main operation is written here. It's modifiable
	result, err := tfplugin_string.InsertAt(str, substr, index)

	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Error executing function: "+err.Error()))
		return
	} else {
		// return the value of above operation by following this format
		resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &result))
	}
}
