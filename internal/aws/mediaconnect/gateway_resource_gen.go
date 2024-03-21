// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_mediaconnect_gateway", gatewayResource)
}

// gatewayResource returns the Terraform awscc_mediaconnect_gateway resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaConnect::Gateway resource.
func gatewayResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: EgressCidrBlocks
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The range of IP addresses that contribute content or initiate output requests for flows communicating with this gateway. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"egress_cidr_blocks": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The range of IP addresses that contribute content or initiate output requests for flows communicating with this gateway. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
			Required:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GatewayArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the gateway.",
		//	  "type": "string"
		//	}
		"gateway_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the gateway.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GatewayState
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The current status of the gateway.",
		//	  "enum": [
		//	    "CREATING",
		//	    "ACTIVE",
		//	    "UPDATING",
		//	    "ERROR",
		//	    "DELETING",
		//	    "DELETED"
		//	  ],
		//	  "type": "string"
		//	}
		"gateway_state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The current status of the gateway.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the gateway. This name can not be modified after the gateway is created.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the gateway. This name can not be modified after the gateway is created.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Networks
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of networks in the gateway.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The network settings for a gateway.",
		//	    "properties": {
		//	      "CidrBlock": {
		//	        "description": "A unique IP address range to use for this network. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "description": "The name of the network. This name is used to reference the network and must be unique among networks in this gateway.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name",
		//	      "CidrBlock"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 4,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"networks": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CidrBlock
					"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A unique IP address range to use for this network. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the network. This name is used to reference the network and must be unique among networks in this gateway.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of networks in the gateway.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 4),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::MediaConnect::Gateway",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::Gateway").WithTerraformTypeName("awscc_mediaconnect_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cidr_block":         "CidrBlock",
		"egress_cidr_blocks": "EgressCidrBlocks",
		"gateway_arn":        "GatewayArn",
		"gateway_state":      "GatewayState",
		"name":               "Name",
		"networks":           "Networks",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
