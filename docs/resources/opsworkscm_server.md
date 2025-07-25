---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_opsworkscm_server Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::OpsWorksCM::Server
---

# awscc_opsworkscm_server (Resource)

Resource Type definition for AWS::OpsWorksCM::Server



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `instance_profile_arn` (String)
- `instance_type` (String)
- `service_role_arn` (String)

### Optional

- `associate_public_ip_address` (Boolean)
- `backup_id` (String)
- `backup_retention_count` (Number)
- `custom_certificate` (String)
- `custom_domain` (String)
- `custom_private_key` (String)
- `disable_automated_backup` (Boolean)
- `engine` (String)
- `engine_attributes` (Attributes List) (see [below for nested schema](#nestedatt--engine_attributes))
- `engine_model` (String)
- `engine_version` (String)
- `key_pair` (String)
- `preferred_backup_window` (String)
- `preferred_maintenance_window` (String)
- `security_group_ids` (List of String)
- `server_name` (String)
- `subnet_ids` (List of String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `endpoint` (String)
- `id` (String) Uniquely identifies the resource.
- `server_id` (String)

<a id="nestedatt--engine_attributes"></a>
### Nested Schema for `engine_attributes`

Optional:

- `name` (String)
- `value` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_opsworkscm_server.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_opsworkscm_server.example "id"
```
