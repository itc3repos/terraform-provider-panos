---
layout: "panos"
page_title: "panos: panos_system_info"
sidebar_current: "docs-panos-datasource-system-info"
description: |-
  Gets system info from the firewall.
---

# panos_system_info

Use this data source to retrieve "show system info" from the NGFW.

All contents of "show system info" are saved to the `info` variable.  In
addition, the version number of PAN-OS encountered is saved to multiple
fields for ease of access.

## Example Usage

```hcl
data "panos_system_info" "example" {}
```

## Attribute Reference

The following attributes are present:

* `info` - a map containing the contents of `show system info`.
* `version_major` - Major version number.
* `version_minor` - Minor version number.
* `version_patch` - Patch version number.
