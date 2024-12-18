---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cassandra_type Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Cassandra::Type
---

# awscc_cassandra_type (Resource)

Resource schema for AWS::Cassandra::Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `fields` (Attributes Set) Field definitions of the User-Defined Type (see [below for nested schema](#nestedatt--fields))
- `keyspace_name` (String) Name of the Keyspace which contains the User-Defined Type.
- `type_name` (String) Name of the User-Defined Type.

### Read-Only

- `direct_parent_types` (Set of String) List of parent User-Defined Types that directly reference the User-Defined Type in their fields.
- `direct_referring_tables` (Set of String) List of Tables that directly reference the User-Defined Type in their columns.
- `id` (String) Uniquely identifies the resource.
- `keyspace_arn` (String) ARN of the Keyspace which contains the User-Defined Type.
- `last_modified_timestamp` (Number) Timestamp of the last time the User-Defined Type's meta data was modified.
- `max_nesting_depth` (Number) Maximum nesting depth of the User-Defined Type across the field types.

<a id="nestedatt--fields"></a>
### Nested Schema for `fields`

Required:

- `field_name` (String)
- `field_type` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_cassandra_type.example "keyspace_name|type_name"
```
