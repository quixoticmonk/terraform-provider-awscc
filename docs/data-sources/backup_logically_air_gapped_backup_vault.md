---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_backup_logically_air_gapped_backup_vault Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Backup::LogicallyAirGappedBackupVault
---

# awscc_backup_logically_air_gapped_backup_vault (Data Source)

Data Source schema for AWS::Backup::LogicallyAirGappedBackupVault



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `access_policy` (String)
- `backup_vault_arn` (String)
- `backup_vault_name` (String)
- `backup_vault_tags` (Map of String)
- `encryption_key_arn` (String)
- `max_retention_days` (Number)
- `min_retention_days` (Number)
- `notifications` (Attributes) (see [below for nested schema](#nestedatt--notifications))
- `vault_state` (String)
- `vault_type` (String)

<a id="nestedatt--notifications"></a>
### Nested Schema for `notifications`

Read-Only:

- `backup_vault_events` (List of String)
- `sns_topic_arn` (String)
