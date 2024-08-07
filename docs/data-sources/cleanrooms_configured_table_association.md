---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cleanrooms_configured_table_association Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CleanRooms::ConfiguredTableAssociation
---

# awscc_cleanrooms_configured_table_association (Data Source)

Data Source schema for AWS::CleanRooms::ConfiguredTableAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `configured_table_association_analysis_rules` (Attributes List) (see [below for nested schema](#nestedatt--configured_table_association_analysis_rules))
- `configured_table_association_identifier` (String)
- `configured_table_identifier` (String)
- `description` (String)
- `membership_identifier` (String)
- `name` (String)
- `role_arn` (String)
- `tags` (Attributes List) An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--configured_table_association_analysis_rules"></a>
### Nested Schema for `configured_table_association_analysis_rules`

Read-Only:

- `policy` (Attributes) (see [below for nested schema](#nestedatt--configured_table_association_analysis_rules--policy))
- `type` (String)

<a id="nestedatt--configured_table_association_analysis_rules--policy"></a>
### Nested Schema for `configured_table_association_analysis_rules.policy`

Read-Only:

- `v1` (Attributes) (see [below for nested schema](#nestedatt--configured_table_association_analysis_rules--policy--v1))

<a id="nestedatt--configured_table_association_analysis_rules--policy--v1"></a>
### Nested Schema for `configured_table_association_analysis_rules.policy.v1`

Read-Only:

- `aggregation` (Attributes) (see [below for nested schema](#nestedatt--configured_table_association_analysis_rules--policy--v1--aggregation))
- `custom` (Attributes) (see [below for nested schema](#nestedatt--configured_table_association_analysis_rules--policy--v1--custom))
- `list` (Attributes) (see [below for nested schema](#nestedatt--configured_table_association_analysis_rules--policy--v1--list))

<a id="nestedatt--configured_table_association_analysis_rules--policy--v1--aggregation"></a>
### Nested Schema for `configured_table_association_analysis_rules.policy.v1.aggregation`

Read-Only:

- `allowed_additional_analyses` (List of String)
- `allowed_result_receivers` (List of String)


<a id="nestedatt--configured_table_association_analysis_rules--policy--v1--custom"></a>
### Nested Schema for `configured_table_association_analysis_rules.policy.v1.custom`

Read-Only:

- `allowed_additional_analyses` (List of String)
- `allowed_result_receivers` (List of String)


<a id="nestedatt--configured_table_association_analysis_rules--policy--v1--list"></a>
### Nested Schema for `configured_table_association_analysis_rules.policy.v1.list`

Read-Only:

- `allowed_additional_analyses` (List of String)
- `allowed_result_receivers` (List of String)





<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
