// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
    "context"
    "net/url"
    "strings"
    "github.com/hashicorp/terraform-plugin-framework/attr"
    "github.com/hashicorp/terraform-plugin-framework/function"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var URLReturnAttributeTypes = map[string]attr.Type{
    "scheme":       types.StringType,
    "opaque":       types.StringType,
    "userinfo":     types.StringType,
    "hostname":     types.StringType,
    "port":         types.StringType,
    "path":         types.StringType,
    "raw_path":     types.StringType,
    "omit_host":    types.BoolType,
    "force_query":  types.BoolType,
    "fragment":     types.StringType,
    "raw_fragment": types.StringType,
}

var _ function.Function = &UrlParseFunction{}

type UrlParseFunction struct{}

// please register this function name at internal/provider/provider.go on your TF providers
func NewUrlParseFunction() function.Function {
    return &UrlParseFunction{}
}

func (f *UrlParseFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
    resp.Name = "url_parse"
}

// You can modify the required parameters on the Definition func
func (f *UrlParseFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
    resp.Definition = function.Definition{
        Summary:     "Parse given URL string",
        Description: "Parse given URL string to object which has Go lang URL types",
        Parameters: []function.Parameter{
            function.StringParameter{
                Name:        "url",
                Description: "An URL to be parsed",
            },
        },
        Return: function.ObjectReturn{
            AttributeTypes: URLReturnAttributeTypes,
        },
    }
}

// You can modify the required parameters and operations on the Run func
func (f *UrlParseFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
    var urlStr string

    // catch the params by using defined variables
    resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &urlStr))
    if resp.Error != nil {
        return
    }

    // main operation is written here. It's modifiable
    parsedUrl, err := url.Parse(urlStr)

    if err != nil {
        resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Error executing function: " + err.Error()))
        return
    }

    hostStr := strings.Split(string(parsedUrl.Host), ":")
    hostname := hostStr[0]

    var port string
    if len(hostStr) > 1 {
        port = hostStr[1]
    } else {
        port = ""
    }

    // return the value of above operation by following this format
    urlObj, diags := types.ObjectValue(
        URLReturnAttributeTypes,
        map[string]attr.Value{
            "scheme":       types.StringValue(parsedUrl.Scheme),
            "opaque":       types.StringValue(parsedUrl.Opaque),
            "userinfo":     types.StringValue(parsedUrl.User.String()),
            "hostname":     types.StringValue(hostname),
            "port":         types.StringValue(port),
            "path":         types.StringValue(parsedUrl.Path),
            "raw_path":     types.StringValue(parsedUrl.RawPath),
            "omit_host":    types.BoolValue(parsedUrl.OmitHost),
            "force_query":  types.BoolValue(parsedUrl.ForceQuery),
            "fragment":     types.StringValue(parsedUrl.Fragment),
            "raw_fragment": types.StringValue(parsedUrl.RawFragment),
        },
    )

    resp.Error = function.FuncErrorFromDiags(ctx, diags)
    if resp.Error != nil {
        return
    }

    resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &urlObj))
}
