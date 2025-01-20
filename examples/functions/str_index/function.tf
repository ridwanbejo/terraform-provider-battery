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
    sample_1 = provider::battery::str_index("cheese", "e")
    sample_2 = provider::battery::str_index("blah blah blah blah ...", "blah")
    sample_3 = provider::battery::str_index("foo bar foo bar foo bar foo bar foo bar ...", "bar")
    sample_4 = provider::battery::str_index("chicken", "ken")
    sample_5 = provider::battery::str_index("chicken", "foo")
  }
}
