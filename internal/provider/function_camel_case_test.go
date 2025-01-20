// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
    "testing"
    "regexp"

    "github.com/hashicorp/terraform-plugin-testing/helper/resource"
    "github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestCamelCaseFunction_string_with_space_only(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::camel_case("lorem ipsum sit dolor amet")
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "loremIpsumSitDolorAmet"),
                ),
            },
        },
    })
}

func TestCamelCaseFunction_string_with_space_only_and_capitalized(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::camel_case("Lorem ipsum sit dolor amet")
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "loremIpsumSitDolorAmet"),
                ),
            },
        },
    })
}

func TestCamelCaseFunction_string_with_dash_only(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::camel_case("lorem-ipsum-sit-dolor-amet")
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "loremIpsumSitDolorAmet"),
                ),
            },
        },
    })
}

func TestCamelCaseFunction_string_with_dash_only_and_capitalized(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::camel_case("Lorem-ipsum-sit-dolor-amet")
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "loremIpsumSitDolorAmet"),
                ),
            },
        },
    })
}

func TestCamelCaseFunction_string_with_underscore_only(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::camel_case("lorem_ipsum_sit_dolor_amet")
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "loremIpsumSitDolorAmet"),
                ),
            },
        },
    })
}

func TestCamelCaseFunction_string_with_underscore_only_and_capitalized(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::camel_case("Lorem_ipsum_sit_dolor_amet")
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "loremIpsumSitDolorAmet"),
                ),
            },
        },
    })
}

func TestCamelCaseFunction_empty_string(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::camel_case("")
                }
                `,
                ExpectError: regexp.MustCompile(`string length minimum is 1.`),
            },
        },
    })
}


func TestCamelCaseFunction_number_as_string(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::camel_case(12345)
                }
                `,
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckOutput("test", "12345"),
                ),
            },
        },
    })
}

func TestCamelCaseFunction_list_as_string(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::camel_case([1, 2, 3])
                }
                `,
                ExpectError: regexp.MustCompile(`Invalid value for "string" parameter: string required.`),
            },
        },
    })
}

func TestCamelCaseFunction_object_as_string(t *testing.T) {
    resource.UnitTest(t, resource.TestCase{
        TerraformVersionChecks: []tfversion.TerraformVersionCheck{
            tfversion.SkipBelow(tfversion.Version1_8_0),
        },
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: `
                output "test" {
                    value = provider::battery::camel_case({foo="bar", name="John Thor"})
                }
                `,
                ExpectError: regexp.MustCompile(`Invalid value for "string" parameter: string required.`),
            },
        },
    })
}
