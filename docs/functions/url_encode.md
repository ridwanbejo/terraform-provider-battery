---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "url_encode function - battery"
subcategory: ""
description: |-
  Encode URL string
---

# function: url_encode

Encode URL string with escaped characters

## Example Usage

```terraform
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
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
url_encode(url string) string
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `url` (String) url to be encoded

