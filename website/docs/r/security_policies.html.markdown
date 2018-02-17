---
layout: "panos"
page_title: "panos: panos_security_policies"
sidebar_current: "docs-panos-resource-security-policies"
description: |-
  Manages security policies.
---

# panos_security_policies

This resource allows you to add/update/delete security policies.

This resource manages the full set of security policies, enforcing both the
contents of individual rules as well as their ordering.  Rules are defined in
a `rule` config block.

For each security policy, there are three styles of profile settings:

* `None` (the default)
* `Group`
* `Profiles`

The Profile Setting is implicitly chosen based on what params are configured
for the security policy.  If you want a Profile Setting of `Group`, then the
`group` param should be set to the desired Group Profile.  If you want a
Profile Setting of `Profiles`, then you will need to specify one or more of
the following params:

* `virus`
* `spyware`
* `vulnerability`
* `url_filtering`
* `file_blocking`
* `wildfire_analysis`
* `data_filtering`

If the `group` param and none of the `Profiles` params are specified, then
the Profile Setting is set to `None`.

## Example Usage

```hcl
resource "panos_security_policies" "example" {
    rule {
        name = "allow bizdev to dmz"
        source_zones = ["bizdev"]
        source_addresses = ["any"]
        source_users = ["any"]
        hip_profiles = ["any"]
        destination_zones = ["dmz"]
        destination_addresses = ["any"]
        applications = ["any"]
        services = ["application-default"]
        categories = ["any"]
        action = "allow"
    }
    rule {
        name = "deny sales to eng"
        source_zones = ["sales"]
        source_addresses = ["any"]
        source_users = ["any"]
        hip_profiles = ["any"]
        destination_zones = ["eng"]
        destination_addresses = ["any"]
        applications = ["any"]
        services = ["application-default"]
        categories = ["any"]
        action = "deny"
    }
}
```

## Argument Reference

The following arguments are supported:

* `vsys` - (Optional) The vsys to put the security policy into (default:
  `vsys1`).
* `rulebase` - (Optional) The rulebase.  For firewalls, there is only the
  `rulebase` value (default), but on Panorama, there is also `pre-rulebase`
  and `post-rulebase`.
* `rule` - The security policy definition (see below).  The security policy
  ordering will match how they appear in the terraform plan file.

The following arguments are valid for each `rule` section:

* `name` - (Required) The security policy name.
* `type` - (Optional) Rule type.  This can be `universal` (default),
  `interzone`, or `intrazone`.
* `description` - (Optional) The description.
* `tags` - (Optional) List of tags for this security rule.
* `source_zones` - (Required) List of source zones.
* `source_addresses` - (Required) List of source addresses.
* `negate_source` - (Optional, bool) If the source should be negated.
* `source_users` - (Required) List of source users.
* `hip_profiles` - (Required) List of HIP profiles.
* `destination_zones` - (Required) List of destination zones.
* `destination_addresses` - (Required) List of destination addresses.
* `negate_destination` - (Optional, bool) If the destination should be negated.
* `applications` - (Required) List of applications.
* `services` - (Required) List of services.
* `categories` - (Required) List of categories.
* `action` - (Optional) Action for the matched traffic.  This can be `allow`
  (default), `deny`, `drop`, `reset-client`, `reset-server`, or `reset-both`.
* `log_setting` - (Optional) Log forwarding profile.
* `log_start` - (Optional, bool) Log the start of the traffic flow.
* `log_end` - (Optional, bool) Log the end of the traffic flow (default: `true`).
* `disabled` - (Optional, bool) Set to `true` to disable this rule.
* `schedule` - (Optional) The security policy schedule.
* `icmp_unreachable` - (Optional) Set to `true` to enable ICMP unreachable.
* `disable_server_response_inspection` - (Optional) Set to `true` to disable
  server response inspection.
* `group` - (Optional) Profile Setting: `Group` - The group profile name.
* `virus` - (Optional) Profile Setting: `Profiles` - The antivirus setting.
* `spyware` - (Optional) Profile Setting: `Profiles` - The anti-spyware
  setting.
* `vulnerability` - (Optional) Profile Setting: `Profiles` - The Vulnerability
  Protection setting.
* `url_filtering` - (Optional) Profile Setting: `Profiles` - The URL filtering
  setting.
* `file_blocking` - (Optional) Profile Setting: `Profiles` - The file blocking
  setting.
* `wildfire_analysis` - (Optional) Profile Setting: `Profiles` - The WildFire
  Analysis setting.
* `data_filtering` - (Optional) Profile Setting: `Profiles` - The Data
  Filtering setting.
