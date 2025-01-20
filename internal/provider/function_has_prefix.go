// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
    "context"
    "strings"

    "github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = &HasPrefixFunction{}

type HasPrefixFunction struct{}

// please register this function name at internal/provider/provider.go on your TF providers
func NewHasPrefixFunction() function.Function {
    return &HasPrefixFunction{}
}

func (f *HasPrefixFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
    resp.Name = "has_prefix"
}

// You can modify the required parameters on the Definition func
func (f *HasPrefixFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
    resp.Definition = function.Definition{
        Summary:     "Check prefix on given string",
        Description: "Check prefix on given string",
        Parameters: []function.Parameter{
            function.StringParameter{
                Name:        "string",
                Description: "string for checking the prefix",
            },
            function.StringParameter{
                Name:        "prefix",
                Description: "prefix for checking the prefix",
            },
        },
        Return: function.BoolReturn{},
    }
}

// You can modify the required parameters and operations on the Run func
func (f *HasPrefixFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
    var str string
    var prefix string

    // catch the params by using defined variables
    resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str, &prefix))
    if resp.Error != nil {
        return
    }

    // main operation is written here. It's modifiable
    result := strings.HasPrefix(str, prefix)

    // return the value of above operation by following this format
    resp.Error = resp.Result.Set(ctx, &result)
}
