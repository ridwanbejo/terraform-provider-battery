// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
    "context"
    "net/url"

    "github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = &UrlEncodeFunction{}

type UrlEncodeFunction struct{}

// please register this function name at internal/provider/provider.go on your TF providers
func NewUrlEncodeFunction() function.Function {
    return &UrlEncodeFunction{}
}

func (f *UrlEncodeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
    resp.Name = "url_encode"
}

// You can modify the required parameters on the Definition func
func (f *UrlEncodeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
    resp.Definition = function.Definition{
        Summary:     "Encode URL string",
        Description: "Encode URL string with escaped characters",
        Parameters: []function.Parameter{
            function.StringParameter{
                Name:        "url",
                Description: "url to be encoded",
            },
        },
        Return: function.StringReturn{},
    }
}

// You can modify the required parameters and operations on the Run func
func (f *UrlEncodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
    var urlStr string

    // catch the params by using defined variables
    resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &urlStr))
    if resp.Error != nil {
        return
    }

    // main operation is written here. It's modifiable
    result := url.QueryEscape(urlStr)

    // return the value of above operation by following this format
    resp.Error = resp.Result.Set(ctx, &result)
}
