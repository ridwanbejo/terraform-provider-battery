// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
    "context"
    "github.com/hashicorp/terraform-plugin-framework/function"
    tfplugin_string "terraform-provider-battery/internal/libs/string"
)

var _ function.Function = &AddPrefixFunction{}

type AddPrefixFunction struct{}

// please register this function name at internal/provider/provider.go on your TF providers
func NewAddPrefixFunction() function.Function {
    return &AddPrefixFunction{}
}

func (f *AddPrefixFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
    resp.Name = "add_prefix"
}

// You can modify the required parameters on the Definition func
func (f *AddPrefixFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
    resp.Definition = function.Definition{
        Summary:     "Add prefix to given string",
        Description: "Add prefix to given string",
        Parameters: []function.Parameter{
            function.StringParameter{
                Name:        "string",
                Description: "string for which prefix will be added to",
            },
            function.StringParameter{
                Name:        "prefix",
                Description: "prefix to be added on given string",
            },
        },
        Return: function.StringReturn{},
    }
}

// You can modify the required parameters and operations on the Run func
func (f *AddPrefixFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
    var str string
    var prefix string

    // catch the params by using defined variables
    resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str, &prefix))
    if resp.Error != nil {
        return
    }

    // main operation is written here. It's modifiable
    result := tfplugin_string.AddPrefix(str, prefix)

    // return the value of above operation by following this format
    resp.Error = resp.Result.Set(ctx, &result)
}
