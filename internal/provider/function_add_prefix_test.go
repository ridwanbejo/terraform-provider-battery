// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAddPrefixFunction_expected_hello(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_prefix("hello", "__test__")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "__test__hello"),
				),
			},
		},
	})
}

func TestAddPrefixFunction_empty_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_prefix("", "__test__")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "__test__"),
				),
			},
		},
	})
}

func TestAddPrefixFunction_empty_prefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_prefix("hello", "")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "hello"),
				),
			},
		},
	})
}

func TestAddPrefixFunction_number_as_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_prefix(12345, "__test__")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "__test__12345"),
				),
			},
		},
	})
}

func TestAddPrefixFunction_number_as_prefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_prefix("hello", 12345)
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "12345hello"),
				),
			},
		},
	})
}

func TestAddPrefixFunction_list_as_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_prefix([1, 2, 3], "__test__")
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "string" parameter: string required.`),
			},
		},
	})
}

func TestAddPrefixFunction_list_as_prefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_prefix("hello", [1, 2, 3])
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "prefix" parameter: string required.`),
			},
		},
	})
}

func TestAddPrefixFunction_object_as_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_prefix({foo="bar", name="John Thor"}, "__test__")
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "string" parameter: string required.`),
			},
		},
	})
}

func TestAddPrefixFunction_object_as_prefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_prefix("hello", {foo="bar", name="John Thor"})
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "prefix" parameter: string required.`),
			},
		},
	})
}
