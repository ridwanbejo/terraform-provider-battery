// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
    "context"
    "strings"

    "github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = &StrCompareFunction{}

type StrCompareFunction struct{}

func NewStrCompareFunction() function.Function {
    return &StrCompareFunction{}
}

func (f *StrCompareFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
    resp.Name = "str_compare"
}

// You can modify the required parameters on the Definition func
func (f *StrCompareFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
    resp.Definition = function.Definition{
        Summary:     "Compare two given string",
        Description: "Compare returns an integer comparing two strings lexicographically. The result will be 0 if a == b, -1 if a < b, and +1 if a > b. It use strings.Count() from Go Standard Library",
        Parameters: []function.Parameter{
            function.StringParameter{
                Name:        "string",
                Description: "string to be compared with the substring",
            },
            function.StringParameter{
                Name:        "substring",
                Description: "substring to be compared with the string",
            },
        },
        Return: function.Int32Return{},
    }
}

// You can modify the required parameters and operations on the Run func
func (f *StrCompareFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
    var str string
    var substr string

    // catch the params by using defined variables
    resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str, &substr))
    if resp.Error != nil {
        return
    }

    // main operation is written here. It's modifiable
    result := strings.Compare(str, substr)

    // return the value of above operation by following this format
    resp.Error = resp.Result.Set(ctx, &result)
}
