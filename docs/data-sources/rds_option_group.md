---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_rds_option_group Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::RDS::OptionGroup
---

# awscc_rds_option_group (Data Source)

Data Source schema for AWS::RDS::OptionGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `engine_name` (String) Specifies the name of the engine that this option group should be associated with.
 Valid Values: 
  +   ``mariadb`` 
  +   ``mysql`` 
  +   ``oracle-ee`` 
  +   ``oracle-ee-cdb`` 
  +   ``oracle-se2`` 
  +   ``oracle-se2-cdb`` 
  +   ``postgres`` 
  +   ``sqlserver-ee`` 
  +   ``sqlserver-se`` 
  +   ``sqlserver-ex`` 
  +   ``sqlserver-web``
- `major_engine_version` (String) Specifies the major version of the engine that this option group should be associated with.
- `option_configurations` (Attributes List) A list of all available options for an option group. (see [below for nested schema](#nestedatt--option_configurations))
- `option_group_description` (String) The description of the option group.
- `option_group_name` (String) The name of the option group to be created.
 Constraints:
  +  Must be 1 to 255 letters, numbers, or hyphens
  +  First character must be a letter
  +  Can't end with a hyphen or contain two consecutive hyphens
  
 Example: ``myoptiongroup``
 If you don't specify a value for ``OptionGroupName`` property, a name is automatically created for the option group.
  This value is stored as a lowercase string.
- `tags` (Attributes List) Tags to assign to the option group. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--option_configurations"></a>
### Nested Schema for `option_configurations`

Read-Only:

- `db_security_group_memberships` (Set of String) A list of DB security groups used for this option.
- `option_name` (String) The configuration of options to include in a group.
- `option_settings` (Attributes List) The option settings to include in an option group. (see [below for nested schema](#nestedatt--option_configurations--option_settings))
- `option_version` (String) The version for the option.
- `port` (Number) The optional port for the option.
- `vpc_security_group_memberships` (Set of String) A list of VPC security group names used for this option.

<a id="nestedatt--option_configurations--option_settings"></a>
### Nested Schema for `option_configurations.option_settings`

Read-Only:

- `name` (String) The name of the option that has settings that you can set.
- `value` (String) The current value of the option setting.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").
- `value` (String) A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").
