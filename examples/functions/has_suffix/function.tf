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
    sample_1 = provider::battery::has_suffix("hello__test__", "__test__")
    sample_2 = provider::battery::has_suffix("_test_world", "__test__")
  }
}
