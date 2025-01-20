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
    sample_1 = provider::battery::title("lorem ipsum sit dolor amet")
    sample_2 = provider::battery::title("lorem-ipsum-sit-dolor-amet")
    sample_3 = provider::battery::title("lorem_ipsum_sit_dolor_amet")
    sample_4 = provider::battery::title("Lorem ipsum sit dolor amet")
    sample_5 = provider::battery::title("Lorem-ipsum-sit-dolor-amet")
    sample_6 = provider::battery::title("Lorem_ipsum_sit_dolor_amet")
  }
}
