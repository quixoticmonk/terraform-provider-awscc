---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_glue_database Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Glue::Database
---

# awscc_glue_database (Resource)

Resource Type definition for AWS::Glue::Database



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `catalog_id` (String) The AWS account ID for the account in which to create the catalog object.
- `database_input` (Attributes) The metadata for the database. (see [below for nested schema](#nestedatt--database_input))

### Optional

- `database_name` (String) The name of the database. For hive compatibility, this is folded to lowercase when it is store.

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--database_input"></a>
### Nested Schema for `database_input`

Optional:

- `create_table_default_permissions` (Attributes List) Creates a set of default permissions on the table for principals. Used by AWS Lake Formation. Not used in the normal course of AWS Glue operations. (see [below for nested schema](#nestedatt--database_input--create_table_default_permissions))
- `description` (String) A description of the database.
- `federated_database` (Attributes) A FederatedDatabase structure that references an entity outside the AWS Glue Data Catalog. (see [below for nested schema](#nestedatt--database_input--federated_database))
- `location_uri` (String) The location of the database (for example, an HDFS path).
- `name` (String) The name of the database. For hive compatibility, this is folded to lowercase when it is stored.
- `parameters` (String) These key-value pairs define parameters and properties of the database.
- `target_database` (Attributes) A DatabaseIdentifier structure that describes a target database for resource linking. (see [below for nested schema](#nestedatt--database_input--target_database))

<a id="nestedatt--database_input--create_table_default_permissions"></a>
### Nested Schema for `database_input.create_table_default_permissions`

Optional:

- `permissions` (List of String) The permissions that are granted to the principal.
- `principal` (Attributes) The principal who is granted permissions. (see [below for nested schema](#nestedatt--database_input--create_table_default_permissions--principal))

<a id="nestedatt--database_input--create_table_default_permissions--principal"></a>
### Nested Schema for `database_input.create_table_default_permissions.principal`

Optional:

- `data_lake_principal_identifier` (String) An identifier for the AWS Lake Formation principal.



<a id="nestedatt--database_input--federated_database"></a>
### Nested Schema for `database_input.federated_database`

Optional:

- `connection_name` (String) The name of the connection to the external metastore.
- `identifier` (String) A unique identifier for the federated database.


<a id="nestedatt--database_input--target_database"></a>
### Nested Schema for `database_input.target_database`

Optional:

- `catalog_id` (String) The ID of the Data Catalog in which the database resides.
- `database_name` (String) The name of the catalog database.
- `region` (String) Region of the target database.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_glue_database.example "database_name"
```
