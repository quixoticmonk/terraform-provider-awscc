
---
page_title: "awscc_iot_certificate_provider Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Use the AWS::IoT::CertificateProvider resource to declare an AWS IoT Certificate Provider.
---

# awscc_iot_certificate_provider (Resource)

Use the AWS::IoT::CertificateProvider resource to declare an AWS IoT Certificate Provider.

## Example Usage

### IoT Certificate Provider with Lambda Integration

Creates an IoT Certificate Provider that uses a Lambda function to handle certificate operations, with the necessary IAM roles and permissions to allow IoT service to create certificates from CSR requests.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# IAM role for Lambda
data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "lambda_permissions" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = ["arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "iot:CreateCertificateFromCsr",
      "iot:DeleteCertificate",
      "iot:DescribeCertificate",
      "iot:RegisterCertificate",
      "iot:UpdateCertificate"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "lambda_role" {
  role_name                   = "iot-certificate-provider-role"
  assume_role_policy_document = data.aws_iam_policy_document.lambda_assume_role.json
  path                        = "/"

  policies = [
    {
      policy_name     = "IoTCertificateProviderPolicy"
      policy_document = data.aws_iam_policy_document.lambda_permissions.json
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Allow IoT service to invoke the Lambda function
resource "awscc_lambda_permission" "allow_iot_invoke" {
  action        = "lambda:InvokeFunction"
  function_name = "iot-certificate-provider"
  principal     = "iot.amazonaws.com"
  source_arn    = "arn:aws:iot:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:*"
}

# Sample Lambda function for certificate provider
resource "aws_lambda_function" "cert_provider" {
  filename      = "lambda_function.zip"
  function_name = "iot-certificate-provider"
  role          = awscc_iam_role.lambda_role.arn
  handler       = "index.handler"
  runtime       = "nodejs18.x"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a sample Lambda function code
resource "local_file" "lambda_function" {
  filename = "index.js"
  content  = <<-EOT
exports.handler = async (event) => {
    console.log('Received event:', JSON.stringify(event));
    return {
        statusCode: 200,
        body: JSON.stringify('Hello from Certificate Provider!')
    };
};
EOT
}

# Create zip file for Lambda
data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = local_file.lambda_function.filename
  output_path = "lambda_function.zip"
}

# IoT Certificate Provider
resource "awscc_iot_certificate_provider" "example" {
  certificate_provider_name      = "example-provider"
  lambda_function_arn            = aws_lambda_function.cert_provider.arn
  account_default_for_operations = ["CreateCertificateFromCsr"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_default_for_operations` (Set of String)
- `lambda_function_arn` (String)

### Optional

- `certificate_provider_name` (String)
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_iot_certificate_provider.example
  id = "certificate_provider_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_iot_certificate_provider.example "certificate_provider_name"
```
