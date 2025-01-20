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
    sample_1 = provider::battery::str_compare("cheese", "cheeze")
    sample_2 = provider::battery::str_compare("cheese", "cheese")
    sample_3 = provider::battery::str_compare("zheese", "cheese")
  }
}
