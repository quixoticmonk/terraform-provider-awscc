---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_vpclattice_resource_configuration Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::VpcLattice::ResourceConfiguration
---

# awscc_vpclattice_resource_configuration (Data Source)

Data Source schema for AWS::VpcLattice::ResourceConfiguration



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `allow_association_to_sharable_service_network` (Boolean)
- `arn` (String)
- `name` (String)
- `port_ranges` (List of String)
- `protocol_type` (String)
- `resource_configuration_auth_type` (String)
- `resource_configuration_definition` (Attributes) (see [below for nested schema](#nestedatt--resource_configuration_definition))
- `resource_configuration_group_id` (String)
- `resource_configuration_id` (String)
- `resource_configuration_type` (String)
- `resource_gateway_id` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--resource_configuration_definition"></a>
### Nested Schema for `resource_configuration_definition`

Read-Only:

- `arn_resource` (String)
- `dns_resource` (Attributes) (see [below for nested schema](#nestedatt--resource_configuration_definition--dns_resource))
- `ip_resource` (String)

<a id="nestedatt--resource_configuration_definition--dns_resource"></a>
### Nested Schema for `resource_configuration_definition.dns_resource`

Read-Only:

- `domain_name` (String)
- `ip_address_type` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)