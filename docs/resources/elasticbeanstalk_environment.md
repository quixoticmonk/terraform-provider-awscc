---
page_title: "awscc_elasticbeanstalk_environment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::ElasticBeanstalk::Environment
---

# awscc_elasticbeanstalk_environment (Resource)

Resource Type definition for AWS::ElasticBeanstalk::Environment

## Example Usage

### Basic usage with Python platform
In this example, we are going to create an AWS Elastic Beanstalk environment using Python platform. Please refer to [Python platform history](https://docs.aws.amazon.com/elasticbeanstalk/latest/platforms/platform-history-python.html) and use the current platform version and solution stack name. Also, please replace `example-aws-elasticbeanstalk-ec2-role` with your existing [instance profile](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/iam-instanceprofile.html).
```terraform
resource "awscc_elasticbeanstalk_application" "example-app" {
  application_name = "example-app"
  description      = "example-app"
}

resource "awscc_elasticbeanstalk_environment" "example-env" {
  application_name    = awscc_elasticbeanstalk_application.example-app.application_name
  solution_stack_name = "64bit Amazon Linux 2023 v4.0.3 running Python 3.11"
  option_settings = [{
    namespace   = "aws:autoscaling:launchconfiguration"
    option_name = "IamInstanceProfile"
    value       = "example-aws-elasticbeanstalk-ec2-role"
  }]
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_name` (String) The name of the application that is associated with this environment.

### Optional

- `cname_prefix` (String) If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.
- `description` (String) Your description for this environment.
- `environment_name` (String) A unique name for the environment.
- `operations_role` (String) The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.
- `option_settings` (Attributes List) Key-value pairs defining configuration options for this environment, such as the instance type. (see [below for nested schema](#nestedatt--option_settings))
- `platform_arn` (String) The Amazon Resource Name (ARN) of the custom platform to use with the environment.
- `solution_stack_name` (String) The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.
- `tags` (Attributes List) Specifies the tags applied to resources in the environment. (see [below for nested schema](#nestedatt--tags))
- `template_name` (String) The name of the Elastic Beanstalk configuration template to use with the environment.
- `tier` (Attributes) Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks. (see [below for nested schema](#nestedatt--tier))
- `version_label` (String) The name of the application version to deploy.

### Read-Only

- `endpoint_url` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--option_settings"></a>
### Nested Schema for `option_settings`

Optional:

- `namespace` (String) A unique namespace that identifies the option's associated AWS resource.
- `option_name` (String) The name of the configuration option.
- `resource_name` (String) A unique resource name for the option setting. Use it for a time–based scaling configuration option.
- `value` (String) The current value for the configuration option.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag.
- `value` (String) The value for the tag.


<a id="nestedatt--tier"></a>
### Nested Schema for `tier`

Optional:

- `name` (String) The name of this environment tier.
- `type` (String) The type of this environment tier.
- `version` (String) The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_elasticbeanstalk_environment.example
  id = "environment_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_elasticbeanstalk_environment.example "environment_name"
```