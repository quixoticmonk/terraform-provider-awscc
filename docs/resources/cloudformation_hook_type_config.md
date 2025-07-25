---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudformation_hook_type_config Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Specifies the configuration data for a registered hook in CloudFormation Registry.
---

# awscc_cloudformation_hook_type_config (Resource)

Specifies the configuration data for a registered hook in CloudFormation Registry.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `configuration` (String) The configuration data for the extension, in this account and region.
- `configuration_alias` (String) An alias by which to refer to this extension configuration data.
- `type_arn` (String) The Amazon Resource Name (ARN) of the type without version number.
- `type_name` (String) The name of the type being registered.

We recommend that type names adhere to the following pattern: company_or_organization::service::type.

### Read-Only

- `configuration_arn` (String) The Amazon Resource Name (ARN) for the configuration data, in this account and region.
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_cloudformation_hook_type_config.example
  id = "configuration_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_cloudformation_hook_type_config.example "configuration_arn"
```
