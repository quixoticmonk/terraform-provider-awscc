
---
page_title: "awscc_memorydb_parameter_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::MemoryDB::ParameterGroup resource creates an Amazon MemoryDB ParameterGroup.
---

# awscc_memorydb_parameter_group (Resource)

The AWS::MemoryDB::ParameterGroup resource creates an Amazon MemoryDB ParameterGroup.

## Example Usage

### MemoryDB Parameter Group Configuration

Creates a MemoryDB parameter group for Redis 7 with custom configuration including active defragmentation and volatile-LRU memory policy settings.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create a MemoryDB parameter group
resource "awscc_memorydb_parameter_group" "example" {
  family               = "memorydb_redis7"
  description          = "Example parameter group using AWSCC provider"
  parameter_group_name = "example-memorydb-pg"

  parameters = jsonencode({
    "activedefrag" : "yes",
    "maxmemory-policy" : "volatile-lru"
  })

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    },
    {
      key   = "Environment"
      value = "Example"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `family` (String) The name of the parameter group family that this parameter group is compatible with.
- `parameter_group_name` (String) The name of the parameter group.

### Optional

- `description` (String) A description of the parameter group.
- `parameters` (String) An map of parameter names and values for the parameter update. You must supply at least one parameter name and value; subsequent arguments are optional.
- `tags` (Attributes Set) An array of key-value pairs to apply to this parameter group. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the parameter group.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key for the tag. May not be null.
- `value` (String) The tag's value. May be null.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_memorydb_parameter_group.example
  id = "parameter_group_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_memorydb_parameter_group.example "parameter_group_name"
```
