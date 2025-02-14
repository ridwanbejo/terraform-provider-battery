// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = &StrIndexFunction{}

type StrIndexFunction struct{}

func NewStrIndexFunction() function.Function {
	return &StrIndexFunction{}
}

func (f *StrIndexFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "str_index"
}

func (f *StrIndexFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Get index of substring first occurence in given string",
		Description: "Given a string and substring. This function will help you to get index for substring first occurence on given string. It use strings.Index() from Go Standard Library",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "string",
				Description: "string for checking the substring index",
			},
			function.StringParameter{
				Name:        "substring",
				Description: "substring to be checked on given string",
			},
		},
		Return: function.Int32Return{},
	}
}

func (f *StrIndexFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var str string
	var substr string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str, &substr))
	if resp.Error != nil {
		return
	}

	result := strings.Index(str, substr)

	resp.Error = resp.Result.Set(ctx, &result)
}
