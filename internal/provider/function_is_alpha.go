// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
    "context"
    "github.com/hashicorp/terraform-plugin-framework/function"

    tfplugin_string "terraform-provider-battery/internal/libs/string"
)

var _ function.Function = &IsAlphaFunction{}

type IsAlphaFunction struct{}

// please register this function name at internal/provider/provider.go on your TF providers
func NewIsAlphaFunction() function.Function {
    return &IsAlphaFunction{}
}

func (f *IsAlphaFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
    resp.Name = "is_alpha"
}

// You can modify the required parameters on the Definition func
func (f *IsAlphaFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
    resp.Definition = function.Definition{
        Summary:     "Check if given string is containing alphabet only",
        Description: "Detect alphabet A-Za-z on given string",
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
func (f *IsAlphaFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
    var str string

    // catch the params by using defined variables
    resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str))
    if resp.Error != nil {
        return
    }

    // main operation is written here. It's modifiable
    result := tfplugin_string.IsAlpha(str)

    // return the value of above operation by following this format
    resp.Error = resp.Result.Set(ctx, &result)
}
