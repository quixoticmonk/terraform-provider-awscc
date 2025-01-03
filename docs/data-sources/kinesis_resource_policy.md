---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_kinesis_resource_policy Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Kinesis::ResourcePolicy
---

# awscc_kinesis_resource_policy (Data Source)

Data Source schema for AWS::Kinesis::ResourcePolicy



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `resource_arn` (String) The ARN of the AWS Kinesis resource to which the policy applies.
- `resource_policy` (String) A policy document containing permissions to add to the specified resource. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.
