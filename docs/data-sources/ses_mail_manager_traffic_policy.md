---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ses_mail_manager_traffic_policy Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SES::MailManagerTrafficPolicy
---

# awscc_ses_mail_manager_traffic_policy (Data Source)

Data Source schema for AWS::SES::MailManagerTrafficPolicy



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `default_action` (String)
- `max_message_size_bytes` (Number)
- `policy_statements` (Attributes List) (see [below for nested schema](#nestedatt--policy_statements))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `traffic_policy_arn` (String)
- `traffic_policy_id` (String)
- `traffic_policy_name` (String)

<a id="nestedatt--policy_statements"></a>
### Nested Schema for `policy_statements`

Read-Only:

- `action` (String)
- `conditions` (Attributes List) (see [below for nested schema](#nestedatt--policy_statements--conditions))

<a id="nestedatt--policy_statements--conditions"></a>
### Nested Schema for `policy_statements.conditions`

Read-Only:

- `boolean_expression` (Attributes) (see [below for nested schema](#nestedatt--policy_statements--conditions--boolean_expression))
- `ip_expression` (Attributes) (see [below for nested schema](#nestedatt--policy_statements--conditions--ip_expression))
- `string_expression` (Attributes) (see [below for nested schema](#nestedatt--policy_statements--conditions--string_expression))
- `tls_expression` (Attributes) (see [below for nested schema](#nestedatt--policy_statements--conditions--tls_expression))

<a id="nestedatt--policy_statements--conditions--boolean_expression"></a>
### Nested Schema for `policy_statements.conditions.boolean_expression`

Read-Only:

- `evaluate` (Attributes) (see [below for nested schema](#nestedatt--policy_statements--conditions--boolean_expression--evaluate))
- `operator` (String)

<a id="nestedatt--policy_statements--conditions--boolean_expression--evaluate"></a>
### Nested Schema for `policy_statements.conditions.boolean_expression.evaluate`

Read-Only:

- `analysis` (Attributes) (see [below for nested schema](#nestedatt--policy_statements--conditions--boolean_expression--evaluate--analysis))

<a id="nestedatt--policy_statements--conditions--boolean_expression--evaluate--analysis"></a>
### Nested Schema for `policy_statements.conditions.boolean_expression.evaluate.analysis`

Read-Only:

- `analyzer` (String)
- `result_field` (String)




<a id="nestedatt--policy_statements--conditions--ip_expression"></a>
### Nested Schema for `policy_statements.conditions.ip_expression`

Read-Only:

- `evaluate` (Attributes) (see [below for nested schema](#nestedatt--policy_statements--conditions--ip_expression--evaluate))
- `operator` (String)
- `values` (List of String)

<a id="nestedatt--policy_statements--conditions--ip_expression--evaluate"></a>
### Nested Schema for `policy_statements.conditions.ip_expression.evaluate`

Read-Only:

- `attribute` (String)



<a id="nestedatt--policy_statements--conditions--string_expression"></a>
### Nested Schema for `policy_statements.conditions.string_expression`

Read-Only:

- `evaluate` (Attributes) (see [below for nested schema](#nestedatt--policy_statements--conditions--string_expression--evaluate))
- `operator` (String)
- `values` (List of String)

<a id="nestedatt--policy_statements--conditions--string_expression--evaluate"></a>
### Nested Schema for `policy_statements.conditions.string_expression.evaluate`

Read-Only:

- `attribute` (String)



<a id="nestedatt--policy_statements--conditions--tls_expression"></a>
### Nested Schema for `policy_statements.conditions.tls_expression`

Read-Only:

- `evaluate` (Attributes) (see [below for nested schema](#nestedatt--policy_statements--conditions--tls_expression--evaluate))
- `operator` (String)
- `value` (String)

<a id="nestedatt--policy_statements--conditions--tls_expression--evaluate"></a>
### Nested Schema for `policy_statements.conditions.tls_expression.evaluate`

Read-Only:

- `attribute` (String)





<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)