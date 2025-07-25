
---
page_title: "awscc_cognito_user_pool_user_to_group_attachment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Cognito::UserPoolUserToGroupAttachment
---

# awscc_cognito_user_pool_user_to_group_attachment (Resource)

Resource Type definition for AWS::Cognito::UserPoolUserToGroupAttachment

## Example Usage

### Cognito User Group Membership

To add a Cognito user to a user pool group, configure the user pool group attachment with the user pool ID, group name, and username of the target user.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create a Cognito User Pool
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"

  auto_verified_attributes = ["email"]
  username_attributes      = ["email"]

  password_policy {
    minimum_length    = 8
    require_lowercase = true
    require_numbers   = true
    require_symbols   = true
    require_uppercase = true
  }

  schema {
    attribute_data_type = "String"
    mutable             = true
    name                = "email"
    required            = true

    string_attribute_constraints {
      max_length = "2048"
      min_length = "0"
    }
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a Cognito User Pool Group
resource "aws_cognito_user_group" "example" {
  name         = "example-group"
  user_pool_id = aws_cognito_user_pool.example.id
  description  = "Example user pool group"
}

# Create a Cognito User
resource "aws_cognito_user" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  username     = "example@example.com"

  attributes = {
    email          = "example@example.com"
    email_verified = "true"
  }
}

# Attach the user to the group
resource "awscc_cognito_user_pool_user_to_group_attachment" "example" {
  group_name   = aws_cognito_user_group.example.name
  user_pool_id = aws_cognito_user_pool.example.id
  username     = aws_cognito_user.example.username
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_name` (String)
- `user_pool_id` (String)
- `username` (String)

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_cognito_user_pool_user_to_group_attachment.example
  id = "user_pool_id|group_name|username"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_cognito_user_pool_user_to_group_attachment.example "user_pool_id|group_name|username"
```
