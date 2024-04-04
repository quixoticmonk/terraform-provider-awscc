---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_vpclattice_service Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::VpcLattice::Service
---

# awscc_vpclattice_service (Data Source)

Data Source schema for AWS::VpcLattice::Service



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `auth_type` (String)
- `certificate_arn` (String)
- `created_at` (String)
- `custom_domain_name` (String)
- `dns_entry` (Attributes) (see [below for nested schema](#nestedatt--dns_entry))
- `last_updated_at` (String)
- `name` (String)
- `service_id` (String)
- `status` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--dns_entry"></a>
### Nested Schema for `dns_entry`

Read-Only:

- `domain_name` (String)
- `hosted_zone_id` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)