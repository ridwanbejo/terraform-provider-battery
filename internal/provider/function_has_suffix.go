// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = &HasSuffixFunction{}

type HasSuffixFunction struct{}

// please register this function name at internal/provider/provider.go on your TF providers
func NewHasSuffixFunction() function.Function {
	return &HasSuffixFunction{}
}

func (f *HasSuffixFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "has_suffix"
}

// You can modify the required parameters on the Definition func
func (f *HasSuffixFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Check suffix on given string",
		Description: "Check suffix on given string",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "string",
				Description: "string for checking the suffix",
			},
			function.StringParameter{
				Name:        "suffix",
				Description: "suffix for checking the suffix",
			},
		},
		Return: function.BoolReturn{},
	}
}

// You can modify the required parameters and operations on the Run func
func (f *HasSuffixFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var str string
	var suffix string

	// catch the params by using defined variables
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str, &suffix))
	if resp.Error != nil {
		return
	}

	// main operation is written here. It's modifiable
	result := strings.HasSuffix(str, suffix)

	// return the value of above operation by following this format
	resp.Error = resp.Result.Set(ctx, &result)
}
