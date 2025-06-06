// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_eip_association", eIPAssociationResource)
}

// eIPAssociationResource returns the Terraform awscc_ec2_eip_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::EIPAssociation resource.
func eIPAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AllocationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The allocation ID. This is required.",
		//	  "type": "string"
		//	}
		"allocation_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The allocation ID. This is required.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EIP
		// CloudFormation resource type schema:
		//
		//	{
		//	  "anyOf": [
		//	    {},
		//	    {}
		//	  ],
		//	  "description": "",
		//	  "type": "string"
		//	}
		"eip": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"eip_association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstanceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the instance. The instance must have exactly one attached network interface. You can specify either the instance ID or the network interface ID, but not both.",
		//	  "type": "string"
		//	}
		"instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the instance. The instance must have exactly one attached network interface. You can specify either the instance ID or the network interface ID, but not both.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkInterfaceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the network interface. If the instance has more than one network interface, you must specify a network interface ID.\n You can specify either the instance ID or the network interface ID, but not both.",
		//	  "type": "string"
		//	}
		"network_interface_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the network interface. If the instance has more than one network interface, you must specify a network interface ID.\n You can specify either the instance ID or the network interface ID, but not both.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PrivateIpAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.",
		//	  "type": "string"
		//	}
		"private_ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Associates an Elastic IP address with an instance or a network interface. Before you can use an Elastic IP address, you must allocate it to your account. For more information about working with Elastic IP addresses, see [Elastic IP address concepts and rules](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#vpc-eip-overview).\n You must specify ``AllocationId`` and either ``InstanceId``, ``NetworkInterfaceId``, or ``PrivateIpAddress``.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::EIPAssociation").WithTerraformTypeName("awscc_ec2_eip_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocation_id":        "AllocationId",
		"eip":                  "EIP",
		"eip_association_id":   "Id",
		"instance_id":          "InstanceId",
		"network_interface_id": "NetworkInterfaceId",
		"private_ip_address":   "PrivateIpAddress",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
