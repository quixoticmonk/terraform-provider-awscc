---
page_title: "awscc_billingconductor_pricing_rule Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  A markup/discount that is defined for a specific set of services that can later be associated with a pricing plan.
---

# awscc_billingconductor_pricing_rule (Resource)

A markup/discount that is defined for a specific set of services that can later be associated with a pricing plan.

## Example Usage

### Pricing rule with global scope
This example to create pricing rule using global scope
```terraform
resource "awscc_billingconductor_pricing_rule" "example" {
  name = "Markup10percent"

  scope               = "GLOBAL"
  type                = "MARKUP"
  modifier_percentage = 10

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }

  ]
}
```

### Pricing rule with service scope
This example create pricing rule using service scope
```terraform
resource "awscc_billingconductor_pricing_rule" "example_service" {
  name = "S3Discount"

  scope               = "SERVICE"
  service             = "AmazonS3"
  type                = "DISCOUNT"
  modifier_percentage = 5

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }

  ]
}
```

### Pricing rule with with tier type
This example enable free-tier
```terraform
resource "awscc_billingconductor_pricing_rule" "example_tiering" {
  name = "EnableFreeTiering"

  scope = "GLOBAL"
  type  = "TIERING"

  tiering = {
    free_tier = {
      activated = true
    }
  }


  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }

  ]
}
```

### Pricing rule with billing entity scope
This example create using billing entity scope to markup when marketplace is in use.
```terraform
resource "awscc_billingconductor_pricing_rule" "example_billing_entity" {
  name = "MarketplaceDiscount"

  scope               = "BILLING_ENTITY"
  billing_entity      = "AWS Marketplace"
  type                = "MARKUP"
  modifier_percentage = 5

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }

  ]
}
```

### Pricing rule with SKU scope
This example provides using SKU as scope to provides discount billing item that uses t2.micro on Linux/Unix in Singapore region
```terraform
resource "awscc_billingconductor_pricing_rule" "example_sku" {
  name        = "DiscountEC2_T2Micro_LinuxUnix"
  description = "5% Discount for t2.micro on Linux/Unix in Singapore region"

  scope      = "SKU"
  service    = "AmazonEC2"
  usage_type = "APS1-BoxUsage:t2.medium"
  operation  = "RunInstances"

  type                = "DISCOUNT"
  modifier_percentage = 5

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }

  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Pricing rule name
- `scope` (String) A term used to categorize the granularity of a Pricing Rule.
- `type` (String) One of MARKUP, DISCOUNT or TIERING that describes the behaviour of the pricing rule.

### Optional

- `billing_entity` (String) The seller of services provided by AWS, their affiliates, or third-party providers selling services via AWS Marketplaces. Supported billing entities are AWS, AWS Marketplace, and AISPL.
- `description` (String) Pricing rule description
- `modifier_percentage` (Number) Pricing rule modifier percentage
- `operation` (String) The Operation which a SKU pricing rule is modifying
- `service` (String) The service which a pricing rule is applied on
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))
- `tiering` (Attributes) The set of tiering configurations for the pricing rule. (see [below for nested schema](#nestedatt--tiering))
- `usage_type` (String) The UsageType which a SKU pricing rule is modifying

### Read-Only

- `arn` (String) Pricing rule ARN
- `associated_pricing_plan_count` (Number) The number of pricing plans associated with pricing rule
- `creation_time` (Number) Creation timestamp in UNIX epoch time format
- `id` (String) Uniquely identifies the resource.
- `last_modified_time` (Number) Latest modified timestamp in UNIX epoch time format

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--tiering"></a>
### Nested Schema for `tiering`

Optional:

- `free_tier` (Attributes) The possible customizable free tier configurations. (see [below for nested schema](#nestedatt--tiering--free_tier))

<a id="nestedatt--tiering--free_tier"></a>
### Nested Schema for `tiering.free_tier`

Optional:

- `activated` (Boolean)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_billingconductor_pricing_rule.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_billingconductor_pricing_rule.example "arn"
```