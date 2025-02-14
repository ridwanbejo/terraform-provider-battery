// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAddSuffixFunction_expected_hello(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_suffix("hello", "__test__")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "hello__test__"),
				),
			},
		},
	})
}

func TestAddSuffixFunction_empty_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_suffix("", "__test__")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "__test__"),
				),
			},
		},
	})
}

func TestAddSuffixFunction_empty_suffix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_suffix("hello", "")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "hello"),
				),
			},
		},
	})
}

func TestAddSuffixFunction_number_as_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_suffix(12345, "__test__")
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "12345__test__"),
				),
			},
		},
	})
}

func TestAddSuffixFunction_number_as_suffix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_suffix("hello", 12345)
                }
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "hello12345"),
				),
			},
		},
	})
}

func TestAddSuffixFunction_list_as_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_suffix([1, 2, 3], "__test__")
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "string" parameter: string required.`),
			},
		},
	})
}

func TestAddSuffixFunction_list_as_suffix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_suffix("hello", [1, 2, 3])
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "suffix" parameter: string required.`),
			},
		},
	})
}

func TestAddSuffixFunction_object_as_string(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_suffix({foo="bar", name="John Thor"}, "__test__")
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "string" parameter: string required.`),
			},
		},
	})
}

func TestAddSuffixFunction_object_as_suffix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::battery::add_suffix("hello", {foo="bar", name="John Thor"})
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "suffix" parameter: string required.`),
			},
		},
	})
}
