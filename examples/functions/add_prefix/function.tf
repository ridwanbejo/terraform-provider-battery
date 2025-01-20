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
    sample_1 = provider::battery::add_prefix("hello", "__test__")
    sample_2 = provider::battery::add_prefix("world", "__test__")
  }
}
