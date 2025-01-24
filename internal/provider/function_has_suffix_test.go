// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
    "testing"
    "regexp"

    "github.com/hashicorp/terraform-plugin-testing/helper/resource"
    "github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestHasSuffixFunction_valid_string_and_suffix(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::has_suffix("__test__hello__test__", "__test__")
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "true"),
                ),
            },
        },
    })
}

func TestHasSuffixFunction_empty_string(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::has_suffix("", "__test__")
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "false"),
                ),
            },
        },
    })
}

func TestHasSuffixFunction_empty_suffix(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::has_suffix("__test__hello__test__", "")
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "true"),
                ),
            },
        },
    })
}

func TestHasSuffixFunction_number_as_string(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::has_suffix(12345, "__test__")
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "false"),
                ),
            },
        },
    })
}

func TestHasSuffixFunction_number_as_suffix(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::has_suffix("__test__hello__test__", 12345)
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "false"),
                ),
            },
        },
    })
}

func TestHasSuffixFunction_list_as_string(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::has_suffix([1, 2, 3], "__test__")
                }
                `,
                ExpectError: regexp.MustCompile(`Invalid value for "string" parameter: string required.`),
            },
        },
    })
}

func TestHasSuffixFunction_list_as_suffix(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::has_suffix("__test__hello", [1, 2, 3])
                }
                `,
                ExpectError: regexp.MustCompile(`Invalid value for "suffix" parameter: string required.`),
            },
        },
    })
}

func TestHasSuffixFunction_object_as_string(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::has_suffix({foo="bar", name="John Thor"}, "__test__")
                }
                `,
                ExpectError: regexp.MustCompile(`Invalid value for "string" parameter: string required.`),
            },
        },
    })
}

func TestHasSuffixFunction_object_as_suffix(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::has_suffix("__test__hello", {foo="bar", name="John Thor"})
                }
                `,
                ExpectError: regexp.MustCompile(`Invalid value for "suffix" parameter: string required.`),
            },
        },
    })
}
