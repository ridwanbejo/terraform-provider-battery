// Replace "source" with your provider URL
terraform {
  required_providers {
    battery = {
      source = "hashicorp.com/edu/battery"
    }
  }
}

// Replace "battery" with your provider name
provider "battery" {}

// Edit your output here
output "examples" {
  value = {
    sample_1 = provider::battery::has_prefix("__test__hello", "__test__")
    sample_2 = provider::battery::has_prefix("_test_world", "__test__")
    sample_3 = provider::battery::has_prefix("", "__test__")
  }
}
