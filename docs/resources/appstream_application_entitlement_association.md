---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_appstream_application_entitlement_association Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::AppStream::ApplicationEntitlementAssociation
---

# awscc_appstream_application_entitlement_association (Resource)

Resource Type definition for AWS::AppStream::ApplicationEntitlementAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_identifier` (String)
- `entitlement_name` (String)
- `stack_name` (String)

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_appstream_application_entitlement_association.example "stack_name|entitlement_name|application_identifier"
```
