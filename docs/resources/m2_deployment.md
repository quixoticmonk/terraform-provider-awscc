---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_m2_deployment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Represents a deployment resource of an AWS Mainframe Modernization (M2) application to a specified environment
---

# awscc_m2_deployment (Resource)

Represents a deployment resource of an AWS Mainframe Modernization (M2) application to a specified environment



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_id` (String) The application ID.
- `application_version` (Number) The version number of the application to deploy
- `environment_id` (String) The environment ID.

### Read-Only

- `deployment_id` (String) The deployment ID.
- `id` (String) Uniquely identifies the resource.
- `status` (String) The status of the deployment.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_m2_deployment.example "application_id"
```