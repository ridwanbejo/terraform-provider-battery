// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
    "context"
    "github.com/hashicorp/terraform-plugin-framework/function"

    tfplugin_string "terraform-provider-battery/internal/libs/string"
)

var _ function.Function = &KebabCaseFunction{}

type KebabCaseFunction struct{}

// please register this function name at internal/provider/provider.go on your TF providers
func NewKebabCaseFunction() function.Function {
    return &KebabCaseFunction{}
}

func (f *KebabCaseFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
    resp.Name = "kebab_case"
}

// You can modify the required parameters on the Definition func
func (f *KebabCaseFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
    resp.Definition = function.Definition{
        Summary:     "Convert given string to Kebab case",
        Description: "Convert given string to Kebab case by using Go Standard Library",
        Parameters: []function.Parameter{
            function.StringParameter{
                Name:        "string",
                Description: "string to be converted to Kebab case",
            },
        },
        Return: function.StringReturn{},
    }
}

// You can modify the required parameters and operations on the Run func
func (f *KebabCaseFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
    var str string

    // catch the params by using defined variables
    resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str))
    if resp.Error != nil {
        return
    }

    // main operation is written here. It's modifiable
    result := tfplugin_string.KebabCase(str)

    // return the value of above operation by following this format
    resp.Error = resp.Result.Set(ctx, &result)
}
