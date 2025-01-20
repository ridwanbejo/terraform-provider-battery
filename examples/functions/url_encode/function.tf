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
    sample_1 = provider::battery::url_encode("my/cool+blog&about,stuff")
    sample_2 = provider::battery::url_encode("Hello World!")
  }
}
