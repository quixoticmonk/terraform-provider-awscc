---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_local_gateway_route_table_vpc_association Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::LocalGatewayRouteTableVPCAssociation
---

# awscc_ec2_local_gateway_route_table_vpc_association (Data Source)

Data Source schema for AWS::EC2::LocalGatewayRouteTableVPCAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **local_gateway_id** (String) The ID of the local gateway.
- **local_gateway_route_table_id** (String) The ID of the local gateway route table.
- **local_gateway_route_table_vpc_association_id** (String) The ID of the association.
- **state** (String) The state of the association.
- **tags** (Attributes Set) (see [below for nested schema](#nestedatt--tags))
- **vpc_id** (String) The ID of the VPC.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)

