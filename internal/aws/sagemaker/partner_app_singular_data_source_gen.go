// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sagemaker_partner_app", partnerAppDataSource)
}

// partnerAppDataSource returns the Terraform awscc_sagemaker_partner_app data source.
// This Terraform data source corresponds to the CloudFormation AWS::SageMaker::PartnerApp resource.
func partnerAppDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A collection of settings that specify the maintenance schedule for the PartnerApp.",
		//	  "properties": {
		//	    "AdminUsers": {
		//	      "description": "A list of users with administrator privileges for the PartnerApp.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "A collection of AdminUsers for the PartnerApp",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "maxItems": 5,
		//	      "minItems": 0,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "Arguments": {
		//	      "additionalProperties": false,
		//	      "description": "A list of arguments to pass to the PartnerApp.",
		//	      "patternProperties": {
		//	        "": {
		//	          "maxLength": 1024,
		//	          "pattern": "",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"application_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AdminUsers
				"admin_users": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of users with administrator privileges for the PartnerApp.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Arguments
				"arguments":         // Pattern: ""
				schema.MapAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of arguments to pass to the PartnerApp.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A collection of settings that specify the maintenance schedule for the PartnerApp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the created PartnerApp.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "arn:aws[a-z\\-]*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:partner-app/app-[A-Z0-9]{12}$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the created PartnerApp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AuthType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Auth type of PartnerApp.",
		//	  "enum": [
		//	    "IAM"
		//	  ],
		//	  "type": "string"
		//	}
		"auth_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Auth type of PartnerApp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BaseUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AppServerUrl based on app and account-info.",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"base_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AppServerUrl based on app and account-info.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClientToken
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The client token for the PartnerApp.",
		//	  "maxLength": 36,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"client_token": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The client token for the PartnerApp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnableIamSessionBasedIdentity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Enables IAM Session based Identity for PartnerApp.",
		//	  "type": "boolean"
		//	}
		"enable_iam_session_based_identity": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Enables IAM Session based Identity for PartnerApp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ExecutionRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The execution role for the user.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$",
		//	  "type": "string"
		//	}
		"execution_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The execution role for the user.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaintenanceConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A collection of settings that specify the maintenance schedule for the PartnerApp.",
		//	  "properties": {
		//	    "MaintenanceWindowStart": {
		//	      "description": "The maintenance window start day and time for the PartnerApp.",
		//	      "maxLength": 9,
		//	      "pattern": "(Mon|Tue|Wed|Thu|Fri|Sat|Sun):([01]\\d|2[0-3]):([0-5]\\d)",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "MaintenanceWindowStart"
		//	  ],
		//	  "type": "object"
		//	}
		"maintenance_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaintenanceWindowStart
				"maintenance_window_start": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The maintenance window start day and time for the PartnerApp.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A collection of settings that specify the maintenance schedule for the PartnerApp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the PartnerApp.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9]+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the PartnerApp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags to apply to the PartnerApp.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of tags to apply to the PartnerApp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tier of the PartnerApp.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"tier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The tier of the PartnerApp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of PartnerApp.",
		//	  "enum": [
		//	    "lakera-guard",
		//	    "comet",
		//	    "deepchecks-llm-evaluation",
		//	    "fiddler"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of PartnerApp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SageMaker::PartnerApp",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::PartnerApp").WithTerraformTypeName("awscc_sagemaker_partner_app")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"admin_users":                       "AdminUsers",
		"application_config":                "ApplicationConfig",
		"arguments":                         "Arguments",
		"arn":                               "Arn",
		"auth_type":                         "AuthType",
		"base_url":                          "BaseUrl",
		"client_token":                      "ClientToken",
		"enable_iam_session_based_identity": "EnableIamSessionBasedIdentity",
		"execution_role_arn":                "ExecutionRoleArn",
		"key":                               "Key",
		"maintenance_config":                "MaintenanceConfig",
		"maintenance_window_start":          "MaintenanceWindowStart",
		"name":                              "Name",
		"tags":                              "Tags",
		"tier":                              "Tier",
		"type":                              "Type",
		"value":                             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}