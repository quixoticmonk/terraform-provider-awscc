---
page_title: "awscc_ec2_network_insights_access_scope_analysis Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::EC2::NetworkInsightsAccessScopeAnalysis
---

# awscc_ec2_network_insights_access_scope_analysis (Resource)

Resource schema for AWS::EC2::NetworkInsightsAccessScopeAnalysis

## Example Usage

```terraform
resource "awscc_ec2_network_insights_access_scope_analysis" "example" {
  network_insights_access_scope_id = awscc_ec2_network_insights_access_scope.example.id

  tags = [{
    key   = "Name"
    value = "analysis1"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_insights_access_scope_id` (String)

### Optional

- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `analyzed_eni_count` (Number)
- `end_date` (String)
- `findings_found` (String)
- `id` (String) Uniquely identifies the resource.
- `network_insights_access_scope_analysis_arn` (String)
- `network_insights_access_scope_analysis_id` (String)
- `start_date` (String)
- `status` (String)
- `status_message` (String)

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
  to = awscc_ec2_network_insights_access_scope_analysis.example
  id = "network_insights_access_scope_analysis_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_network_insights_access_scope_analysis.example "network_insights_access_scope_analysis_id"
```