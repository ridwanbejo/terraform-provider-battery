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
    sample_1 = provider::battery::words_count("lorem ipsum sit dolor amet")
    sample_2 = provider::battery::words_count("lorem ipsum sit dolor amet lorem ipsum sit dolor amet")
    sample_3 = provider::battery::words_count("lorem ipsum sit dolor amet lorem ipsum sit dolor amet lorem ipsum sit dolor amet")
  }
}
