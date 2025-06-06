---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_datazone_project_profile Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::DataZone::ProjectProfile Resource Type
---

# awscc_datazone_project_profile (Resource)

Definition of AWS::DataZone::ProjectProfile Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `description` (String)
- `domain_identifier` (String)
- `domain_unit_identifier` (String)
- `status` (String)

### Read-Only

- `created_at` (String)
- `created_by` (String)
- `domain_id` (String)
- `domain_unit_id` (String)
- `id` (String) Uniquely identifies the resource.
- `identifier` (String)
- `last_updated_at` (String)
- `project_profile_id` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_datazone_project_profile.example "domain_identifier|identifier"
```
