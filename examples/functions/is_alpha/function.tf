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
    sample_1 = provider::battery::is_alpha("hello")
    sample_2 = provider::battery::is_alpha("hello-world")
    sample_3 = provider::battery::is_alpha("helloworld")
    sample_4 = provider::battery::is_alpha("hello world!")
    sample_5 = provider::battery::is_alpha("hello_world")
    sample_6 = provider::battery::is_alpha("12345")
    sample_7 = provider::battery::is_alpha("hello12345")
    sample_8 = provider::battery::is_alpha("hello 12345")
    sample_9 = provider::battery::is_alpha("hello 12345!")
  }
}
