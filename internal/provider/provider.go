// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &BatteryProvider{}
var _ provider.ProviderWithFunctions = &BatteryProvider{}

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
    return func() provider.Provider {
        return &BatteryProvider{
            version: version,
        }
    }
}

// BatteryProvider is the provider implementation.
type BatteryProvider struct {
    // version is set to the provider version on release, "dev" when the
    // provider is built and ran locally, and "test" when running acceptance
    // testing.
    version string
}

// Metadata returns the provider type name.
func (p *BatteryProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
    resp.TypeName = "battery"
    resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *BatteryProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
    resp.Schema = schema.Schema{}
}

// Configure prepares an API client for data sources and resources.
func (p *BatteryProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *BatteryProvider) DataSources(_ context.Context) []func() datasource.DataSource {
    return nil
}

// Resources defines the resources implemented in the provider.
func (p *BatteryProvider) Resources(_ context.Context) []func() resource.Resource {
    return nil
}

// Functions defines the functions implemented in the provider.
func (p *BatteryProvider) Functions(_ context.Context) []func() function.Function {
    return []func() function.Function{
        NewStrCountFunction,
        NewStrIndexFunction,
        NewStrCompareFunction,
        NewPascalCaseFunction,
        NewCamelCaseFunction,
        NewTitleFunction,
        NewSnakeCaseFunction,
        NewCapitalizeFunction,
        NewWordsCountFunction,
        NewAddPrefixFunction,
        NewAddSuffixFunction,
        NewHasPrefixFunction,
        NewHasSuffixFunction,
        NewUrlEncodeFunction,
        NewUrlDecodeFunction,
        NewIsAlphaFunction,
        NewIsAlnumFunction,
        NewIsNumericFunction,
        NewKebabCaseFunction,
        NewStrInsertFunction,
        NewUrlParseFunction,
    }
}

