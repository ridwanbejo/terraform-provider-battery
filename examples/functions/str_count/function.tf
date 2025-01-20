terraform {
  required_providers {
    battery = {
      source = "hashicorp.com/edu/battery"
    }
  }
}

provider "battery" {}

output "examples" {
  value = {
    sample_1 = provider::battery::str_count("cheese", "e")
    sample_2 = provider::battery::str_count("blah blah blah blah ...", "blah")
    sample_3 = provider::battery::str_count("foo bar foo bar foo bar foo bar foo bar ...", "foo")
  }
}
