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
    sample_1 = provider::battery::str_insert("Hello world!", "foo", 0)
    sample_2 = provider::battery::str_insert("Hello world!", "foo", 1)
    sample_3 = provider::battery::str_insert("Hello world!", "foo", 2)
    sample_4 = provider::battery::str_insert("Hello world!", "foo", 3)
    sample_5 = provider::battery::str_insert("Hello world!", "foo", 4)
    sample_6 = provider::battery::str_insert("Hello world!", "foo", 5)
    sample_7 = provider::battery::str_insert("Hello world!", "foo", 10)
    sample_9 = provider::battery::str_insert("Hello world!", "foo", 11)
    sample_9 = provider::battery::str_insert("Hello world!", "foo", 12)
    # sample_10 = provider::battery::str_insert("Hello world!", "foo", -1) ---> Throw Error
    # sample_11 = provider::battery::str_insert("Hello world!", "foo", 13) ---> Throw Error
  }
}
