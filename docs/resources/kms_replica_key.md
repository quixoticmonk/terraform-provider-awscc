---
page_title: "awscc_kms_replica_key Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::KMS::ReplicaKey resource specifies a multi-region replica AWS KMS key in AWS Key Management Service (AWS KMS).
---

# awscc_kms_replica_key (Resource)

The AWS::KMS::ReplicaKey resource specifies a multi-region replica AWS KMS key in AWS Key Management Service (AWS KMS).

## Example Usage

### Create replica key in us-east-2 with the primary in us-east-1.

```terraform
provider "awscc" {
  region = "us-east-1"
}

provider "awscc" {
  region = "us-east-2"
  alias  = "secondary"
}

resource "awscc_kms_replica_key" "example" {
  provider = awscc.secondary

  primary_key_arn        = awscc_kms_key.example.arn
  description            = "Example KMS replica key"
  enabled                = true
  pending_window_in_days = 7
  key_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Id" : "key_policy",
      "Statement" : [
        {
          "Sid" : "Enable IAM User Permissions",
          "Effect" : "Allow",
          "Principal" : {
            "AWS" : "arn:aws:iam::0123456789012:root"
          },
          "Action" : "kms:*",
          "Resource" : "*"
        },
        {
          "Sid" : "Allow administration of the key",
          "Effect" : "Allow",
          "Principal" : {
            "AWS" : "arn:aws:iam::0123456789012:role/Admin"
          },
          "Action" : [
            "kms:Create*",
            "kms:Delete*",
            "kms:Disable*",
            "kms:Describe*",
            "kms:Enable*",
            "kms:Get*",
            "kms:List*",
            "kms:Put*",
            "kms:Revoke*",
            "kms:UpdateAlias",
            "kms:ScheduleKeyDeletion",
            "kms:CancelKeyDeletion"
          ],
          "Resource" : "*"
        }
      ]
    }
  )
}

resource "awscc_kms_key" "example" {
  provider     = awscc
  description  = "multi region primary key"
  multi_region = true
  key_policy = jsonencode({
    "Version" : "2012-10-17",
    "Id" : "example_policy",
    "Statement" : [
      {
        "Sid" : "Enable IAM User Permissions",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:aws:iam::0123456789012:root"
        },
        "Action" : "kms:*",
        "Resource" : "*"
      },
    ],
    }
  )
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `key_policy` (String) The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.
- `primary_key_arn` (String) Identifies the primary AWS KMS key to create a replica of. Specify the Amazon Resource Name (ARN) of the AWS KMS key. You cannot specify an alias or key ID. For help finding the ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.

### Optional

- `description` (String) A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.
- `enabled` (Boolean) Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.
- `pending_window_in_days` (Number) Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `key_id` (String)

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
  to = awscc_kms_replica_key.example
  id = "key_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_kms_replica_key.example "key_id"
```