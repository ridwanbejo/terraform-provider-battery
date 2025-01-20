// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
    "context"
    "strings"

    "github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = &StrCountFunction{}

type StrCountFunction struct{}

func NewStrCountFunction() function.Function {
    return &StrCountFunction{}
}

func (f *StrCountFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
    resp.Name = "str_count"
}

func (f *StrCountFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
    resp.Definition = function.Definition{
        Summary:     "Count of substring occurences in given string",
        Description: "Given a string and substring. This function will help you to check substring occurences on given string as superset. It use strings.Count() from Go Standard Library",
        Parameters: []function.Parameter{
            function.StringParameter{
                Name:        "string",
                Description: "string for checking the substring",
            },
            function.StringParameter{
                Name:        "substring",
                Description: "substring to be checked on given string",
            },
        },
        Return: function.Int32Return{},
    }
}

func (f *StrCountFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
    var str string
    var substr string

    resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str, &substr))
    if resp.Error != nil {
        return
    }

    result := strings.Count(str, substr)

    resp.Error = resp.Result.Set(ctx, &result)
}
