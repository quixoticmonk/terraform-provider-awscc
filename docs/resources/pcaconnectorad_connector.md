---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_pcaconnectorad_connector Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::PCAConnectorAD::Connector Resource Type
---

# awscc_pcaconnectorad_connector (Resource)

Definition of AWS::PCAConnectorAD::Connector Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `certificate_authority_arn` (String)
- `directory_id` (String)
- `vpc_information` (Attributes) (see [below for nested schema](#nestedatt--vpc_information))

### Optional

- `tags` (Map of String)

### Read-Only

- `connector_arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--vpc_information"></a>
### Nested Schema for `vpc_information`

Required:

- `security_group_ids` (List of String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_pcaconnectorad_connector.example "connector_arn"
```
