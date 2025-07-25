---
page_title: "awscc_appconfig_environment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::AppConfig::Environment
---

# awscc_appconfig_environment (Resource)

Resource Type definition for AWS::AppConfig::Environment

## Example Usage

```terraform
resource "awscc_appconfig_environment" "example" {
  application_id = awscc_appconfig_application.example.application_id
  name           = "example"
  description    = "Example environment"


  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_id` (String) The application ID.
- `name` (String) A name for the environment.

### Optional

- `deletion_protection_check` (String) On resource deletion this controls whether the Deletion Protection check should be applied, bypassed, or (the default) whether the behavior should be controlled by the account-level Deletion Protection setting. See https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html
- `description` (String) A description of the environment.
- `monitors` (Attributes List) Amazon CloudWatch alarms to monitor during the deployment process. (see [below for nested schema](#nestedatt--monitors))
- `tags` (Attributes Set) Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `environment_id` (String) The environment ID.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--monitors"></a>
### Nested Schema for `monitors`

Optional:

- `alarm_arn` (String) Amazon Resource Name (ARN) of the Amazon CloudWatch alarm.
- `alarm_role_arn` (String) ARN of an AWS Identity and Access Management (IAM) role for AWS AppConfig to monitor AlarmArn.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key-value string map. The valid character set is [a-zA-Z1-9+-=._:/]. The tag key can be up to 128 characters and must not start with aws:.
- `value` (String) The tag value can be up to 256 characters.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_appconfig_environment.example
  id = "application_id|environment_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_appconfig_environment.example "application_id|environment_id"
```