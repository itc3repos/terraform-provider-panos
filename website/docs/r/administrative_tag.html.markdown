---
layout: "panos"
page_title: "panos: panos_administrative_tag"
sidebar_current: "docs-panos-resource-administrative-tag"
description: |-
  Manages administrative tags.
---

# panos_administrative_tag

This resource allows you to add/update/delete administrative tags.

## Example Usage

```hcl
resource "panos_administrative_tag" "example" {
    name = "tag1"
    vsys = "vsys2"
    color = "color5"
    comment = "Internal resources"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The administrative tag's name.
* `vsys` - (Optional) The vsys to put the administrative tag into (default:
  `vsys1`).
* `color` - (Optional) The tag's color.  This should be either an empty string
  (no color) or a string such as `color1` or `color15`.  Note that for maximum
  portability, you should limit color usage to `color16`, which was available
  in PAN-OS 6.1.  PAN-OS 8.1's colors go up to `color42`.  The value `color18`
  is reserved internally by PAN-OS and thus not available for use.
* `comment` - (Optional) The administrative tag's description.
