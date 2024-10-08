---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_appstream_entitlement Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::AppStream::Entitlement
---

# awscc_appstream_entitlement (Resource)

Resource Type definition for AWS::AppStream::Entitlement



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_visibility` (String)
- `attributes` (Attributes Set) (see [below for nested schema](#nestedatt--attributes))
- `name` (String)
- `stack_name` (String)

### Optional

- `description` (String)

### Read-Only

- `created_time` (String)
- `id` (String) Uniquely identifies the resource.
- `last_modified_time` (String)

<a id="nestedatt--attributes"></a>
### Nested Schema for `attributes`

Required:

- `name` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_appstream_entitlement.example "stack_name|name"
```
