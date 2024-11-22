// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_iotsitewise_portal", portalResource)
}

// portalResource returns the Terraform awscc_iotsitewise_portal resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTSiteWise::Portal resource.
func portalResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Alarms
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal. You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range.",
		//	  "properties": {
		//	    "AlarmRoleArn": {
		//	      "description": "The ARN of the IAM role that allows the alarm to perform actions and access AWS resources and services, such as AWS IoT Events.",
		//	      "type": "string"
		//	    },
		//	    "NotificationLambdaArn": {
		//	      "description": "The ARN of the AWS Lambda function that manages alarm notifications. For more information, see Managing alarm notifications in the AWS IoT Events Developer Guide.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"alarms": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AlarmRoleArn
				"alarm_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the IAM role that allows the alarm to perform actions and access AWS resources and services, such as AWS IoT Events.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: NotificationLambdaArn
				"notification_lambda_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the AWS Lambda function that manages alarm notifications. For more information, see Managing alarm notifications in the AWS IoT Events Developer Guide.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal. You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NotificationSenderEmail
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The email address that sends alarm notifications.",
		//	  "type": "string"
		//	}
		"notification_sender_email": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The email address that sends alarm notifications.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PortalArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the portal, which has the following format.",
		//	  "type": "string"
		//	}
		"portal_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the portal, which has the following format.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PortalAuthMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The service to use to authenticate users to the portal. Choose from SSO or IAM. You can't change this value after you create a portal.",
		//	  "type": "string"
		//	}
		"portal_auth_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The service to use to authenticate users to the portal. Choose from SSO or IAM. You can't change this value after you create a portal.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PortalClientId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS SSO application generated client ID (used with AWS SSO APIs).",
		//	  "type": "string"
		//	}
		"portal_client_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS SSO application generated client ID (used with AWS SSO APIs).",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PortalContactEmail
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS administrator's contact email address.",
		//	  "type": "string"
		//	}
		"portal_contact_email": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS administrator's contact email address.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: PortalDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the portal.",
		//	  "type": "string"
		//	}
		"portal_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the portal.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PortalId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the portal.",
		//	  "type": "string"
		//	}
		"portal_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the portal.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PortalName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A friendly name for the portal.",
		//	  "type": "string"
		//	}
		"portal_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A friendly name for the portal.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: PortalStartUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The public root URL for the AWS IoT AWS IoT SiteWise Monitor application portal.",
		//	  "type": "string"
		//	}
		"portal_start_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The public root URL for the AWS IoT AWS IoT SiteWise Monitor application portal.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PortalType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of portal",
		//	  "enum": [
		//	    "SITEWISE_PORTAL_V1",
		//	    "SITEWISE_PORTAL_V2"
		//	  ],
		//	  "type": "string"
		//	}
		"portal_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of portal",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SITEWISE_PORTAL_V1",
					"SITEWISE_PORTAL_V2",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PortalTypeConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Map to associate detail of configuration related with a PortalType.",
		//	  "patternProperties": {
		//	    "": {
		//	      "additionalProperties": false,
		//	      "description": "Container associated a certain PortalType.",
		//	      "properties": {
		//	        "PortalTools": {
		//	          "description": "List of enabled Tools for a certain portal.",
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "PortalTools"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"portal_type_configuration": // Pattern: ""
		schema.MapNestedAttribute{   /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: PortalTools
					"portal_tools": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "List of enabled Tools for a certain portal.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Map to associate detail of configuration related with a PortalType.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf.",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the portal.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted.",
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
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
			Description: "A list of key-value pairs that contain metadata for the portal.",
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
		Description: "Resource schema for AWS::IoTSiteWise::Portal",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Portal").WithTerraformTypeName("awscc_iotsitewise_portal")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alarm_role_arn":            "AlarmRoleArn",
		"alarms":                    "Alarms",
		"key":                       "Key",
		"notification_lambda_arn":   "NotificationLambdaArn",
		"notification_sender_email": "NotificationSenderEmail",
		"portal_arn":                "PortalArn",
		"portal_auth_mode":          "PortalAuthMode",
		"portal_client_id":          "PortalClientId",
		"portal_contact_email":      "PortalContactEmail",
		"portal_description":        "PortalDescription",
		"portal_id":                 "PortalId",
		"portal_name":               "PortalName",
		"portal_start_url":          "PortalStartUrl",
		"portal_tools":              "PortalTools",
		"portal_type":               "PortalType",
		"portal_type_configuration": "PortalTypeConfiguration",
		"role_arn":                  "RoleArn",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
