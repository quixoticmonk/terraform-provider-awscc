---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_agent_status Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Connect::AgentStatus
---

# awscc_connect_agent_status (Data Source)

Data Source schema for AWS::Connect::AgentStatus



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `agent_status_arn` (String) The Amazon Resource Name (ARN) of the agent status.
- `description` (String) The description of the status.
- `display_order` (Number) The display order of the status.
- `instance_arn` (String) The identifier of the Amazon Connect instance.
- `last_modified_region` (String) Last modified region.
- `last_modified_time` (Number) Last modified time.
- `name` (String) The name of the status.
- `reset_order_number` (Boolean) A number indicating the reset order of the agent status.
- `state` (String) The state of the status.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `type` (String) The type of agent status.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
