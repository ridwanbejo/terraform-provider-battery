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
    sample_1 = provider::battery::url_decode("my%2Fcool%2Bblog%26about%2Cstuff")
    sample_2 = provider::battery::url_decode("Hello+World%21")
  }
}
