---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_servicecatalogappregistry_application Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ServiceCatalogAppRegistry::Application
---

# awscc_servicecatalogappregistry_application (Data Source)

Data Source schema for AWS::ServiceCatalogAppRegistry::Application



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `application_id` (String)
- `application_name` (String) The name of the application.
- `application_tag_key` (String) The key of the AWS application tag, which is awsApplication. Applications created before 11/13/2023 or applications without the AWS application tag resource group return no value.
- `application_tag_value` (String) The value of the AWS application tag, which is the identifier of an associated resource. Applications created before 11/13/2023 or applications without the AWS application tag resource group return no value.
- `arn` (String)
- `description` (String) The description of the application.
- `name` (String) The name of the application.
- `tags` (Map of String)
