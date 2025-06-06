---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_ipam Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::IPAM
---

# awscc_ec2_ipam (Data Source)

Data Source schema for AWS::EC2::IPAM



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the IPAM.
- `default_resource_discovery_association_id` (String) The Id of the default association to the default resource discovery, created with this IPAM.
- `default_resource_discovery_id` (String) The Id of the default resource discovery, created with this IPAM.
- `default_resource_discovery_organizational_unit_exclusions` (Attributes Set) A set of organizational unit (OU) exclusions for the default resource discovery, created with this IPAM. (see [below for nested schema](#nestedatt--default_resource_discovery_organizational_unit_exclusions))
- `description` (String)
- `enable_private_gua` (Boolean) Enable provisioning of GUA space in private pools.
- `ipam_id` (String) Id of the IPAM.
- `metered_account` (String) A metered account is an account that is charged for active IP addresses managed in IPAM
- `operating_regions` (Attributes Set) The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring (see [below for nested schema](#nestedatt--operating_regions))
- `private_default_scope_id` (String) The Id of the default scope for publicly routable IP space, created with this IPAM.
- `public_default_scope_id` (String) The Id of the default scope for publicly routable IP space, created with this IPAM.
- `resource_discovery_association_count` (Number) The count of resource discoveries associated with this IPAM.
- `scope_count` (Number) The number of scopes that currently exist in this IPAM.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `tier` (String) The tier of the IPAM.

<a id="nestedatt--default_resource_discovery_organizational_unit_exclusions"></a>
### Nested Schema for `default_resource_discovery_organizational_unit_exclusions`

Read-Only:

- `organizations_entity_path` (String) An AWS Organizations entity path. Build the path for the OU(s) using AWS Organizations IDs separated by a '/'. Include all child OUs by ending the path with '/*'.


<a id="nestedatt--operating_regions"></a>
### Nested Schema for `operating_regions`

Read-Only:

- `region_name` (String) The name of the region.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
