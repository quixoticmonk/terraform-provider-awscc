// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_host", hostResource)
}

// hostResource returns the Terraform awscc_ec2_host resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::Host resource.
func hostResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Outpost hardware asset.",
		//	  "type": "string"
		//	}
		"asset_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Outpost hardware asset.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AutoPlacement
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.",
		//	  "type": "string"
		//	}
		"auto_placement": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Availability Zone in which to allocate the Dedicated Host.",
		//	  "type": "string"
		//	}
		"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Availability Zone in which to allocate the Dedicated Host.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HostId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ID of the host created.",
		//	  "type": "string"
		//	}
		"host_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ID of the host created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HostMaintenance
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Automatically allocates a new dedicated host and moves your instances on to it if a degradation is detected on your current host.",
		//	  "type": "string"
		//	}
		"host_maintenance": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Automatically allocates a new dedicated host and moves your instances on to it if a degradation is detected on your current host.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HostRecovery
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.",
		//	  "type": "string"
		//	}
		"host_recovery": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstanceFamily
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family.",
		//	  "type": "string"
		//	}
		"instance_family": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstanceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.",
		//	  "type": "string"
		//	}
		"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OutpostArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.",
		//	  "type": "string"
		//	}
		"outpost_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Any tags assigned to the Host.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Any tags assigned to the Host.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::EC2::Host",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::Host").WithTerraformTypeName("awscc_ec2_host")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"asset_id":          "AssetId",
		"auto_placement":    "AutoPlacement",
		"availability_zone": "AvailabilityZone",
		"host_id":           "HostId",
		"host_maintenance":  "HostMaintenance",
		"host_recovery":     "HostRecovery",
		"instance_family":   "InstanceFamily",
		"instance_type":     "InstanceType",
		"key":               "Key",
		"outpost_arn":       "OutpostArn",
		"tags":              "Tags",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
