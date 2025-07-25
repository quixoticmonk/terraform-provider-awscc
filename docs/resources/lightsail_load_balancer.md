---
page_title: "awscc_lightsail_load_balancer Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Lightsail::LoadBalancer
---

# awscc_lightsail_load_balancer (Resource)

Resource Type definition for AWS::Lightsail::LoadBalancer

## Example Usage

### Basic Usage

```terraform
resource "awscc_lightsail_instance" "example" {
  blueprint_id  = "nginx"
  bundle_id     = "nano_3_0"
  instance_name = "example-instance"
}

resource "awscc_lightsail_load_balancer" "example" {
  instance_port      = 80
  load_balancer_name = "example-lb"
  attached_instances = [awscc_lightsail_instance.example.instance_name]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

### High Availability With Session Persistence

```terraform
resource "awscc_lightsail_instance" "example_1a" {
  blueprint_id      = "nginx"
  bundle_id         = "nano_3_0"
  instance_name     = "example-instance-1a"
  availability_zone = "us-east-1a"
}

resource "awscc_lightsail_instance" "example_1b" {
  blueprint_id      = "nginx"
  bundle_id         = "nano_3_0"
  instance_name     = "example-instance-1b"
  availability_zone = "us-east-1b"
}

resource "awscc_lightsail_load_balancer" "example" {
  instance_port      = 80
  load_balancer_name = "example-lb"
  attached_instances = [
    awscc_lightsail_instance.example_1a.instance_name,
    awscc_lightsail_instance.example_1b.instance_name
  ]
  session_stickiness_enabled                    = true
  session_stickiness_lb_cookie_duration_seconds = 86500
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `instance_port` (Number) The instance port where you're creating your load balancer.
- `load_balancer_name` (String) The name of your load balancer.

### Optional

- `attached_instances` (Set of String) The names of the instances attached to the load balancer.
- `health_check_path` (String) The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., "/").
- `ip_address_type` (String) The IP address type for the load balancer. The possible values are ipv4 for IPv4 only, and dualstack for IPv4 and IPv6. The default value is dualstack.
- `session_stickiness_enabled` (Boolean) Configuration option to enable session stickiness.
- `session_stickiness_lb_cookie_duration_seconds` (String) Configuration option to adjust session stickiness cookie duration parameter.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `tls_policy_name` (String) The name of the TLS policy to apply to the load balancer.

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `load_balancer_arn` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_lightsail_load_balancer.example
  id = "load_balancer_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_lightsail_load_balancer.example "load_balancer_name"
```