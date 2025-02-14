// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestHasPrefixFunction_valid_string_and_prefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::has_prefix("__test__hello", "__test__")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestHasPrefixFunction_empty_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::has_prefix("", "__test__")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
		},
	})
}

func TestHasPrefixFunction_empty_prefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::has_prefix("__test__hello", "")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "true"),
				),
			},
		},
	})
}

func TestHasPrefixFunction_number_as_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::has_prefix(12345, "__test__")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
		},
	})
}

func TestHasPrefixFunction_number_as_prefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::has_prefix("__test__hello", 12345)
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "false"),
				),
			},
		},
	})
}

func TestHasPrefixFunction_list_as_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::has_prefix([1, 2, 3], "__test__")
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "string" parameter: string required.`),
			},
		},
	})
}

func TestHasPrefixFunction_list_as_prefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::has_prefix("__test__hello", [1, 2, 3])
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "prefix" parameter: string required.`),
			},
		},
	})
}

func TestHasPrefixFunction_object_as_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::has_prefix({foo="bar", name="John Thor"}, "__test__")
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "string" parameter: string required.`),
			},
		},
	})
}

func TestHasPrefixFunction_object_as_prefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::has_prefix("__test__hello", {foo="bar", name="John Thor"})
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "prefix" parameter: string required.`),
			},
		},
	})
}
