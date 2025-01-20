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
    sample_1 = provider::battery::url_parse("http://example.com:80/hello?name=john wick&age=27")
    sample_2 = provider::battery::url_parse("https://example.com:8000/user/1/edit")
    sample_3 = provider::battery::url_parse("https://real-example.com/user/1/edit")
    sample_4 = provider::battery::url_parse("https://fake_user@mail.com:secret12345@real-example.com/user/1/edit")
    sample_5 = provider::battery::url_parse("https://www.mywebsite.com/user/1/edit")
  }
}
