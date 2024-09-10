// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ssmquicksetup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ssmquicksetup_configuration_manager", configurationManagerDataSource)
}

// configurationManagerDataSource returns the Terraform awscc_ssmquicksetup_configuration_manager data source.
// This Terraform data source corresponds to the CloudFormation AWS::SSMQuickSetup::ConfigurationManager resource.
func configurationManagerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ConfigurationDefinitions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "LocalDeploymentAdministrationRoleArn": {
		//	        "type": "string"
		//	      },
		//	      "LocalDeploymentExecutionRoleName": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Parameters": {
		//	        "additionalProperties": false,
		//	        "patternProperties": {
		//	          "": {
		//	            "maxLength": 40960,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Type": {
		//	        "pattern": "^[a-zA-Z0-9_\\-.:/]{3,200}$",
		//	        "type": "string"
		//	      },
		//	      "TypeVersion": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "id": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Parameters",
		//	      "Type"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"configuration_definitions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: LocalDeploymentAdministrationRoleArn
					"local_deployment_administration_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: LocalDeploymentExecutionRoleName
					"local_deployment_execution_role_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Parameters
					"parameters":        // Pattern: ""
					schema.MapAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: TypeVersion
					"type_version": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^.{0,512}$",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"last_modified_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ManagerArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"manager_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^[ A-Za-z0-9_-]{1,50}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: StatusSummaries
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "LastUpdatedAt": {
		//	        "type": "string"
		//	      },
		//	      "Status": {
		//	        "enum": [
		//	          "INITIALIZING",
		//	          "DEPLOYING",
		//	          "SUCCEEDED",
		//	          "DELETING",
		//	          "STOPPING",
		//	          "FAILED",
		//	          "STOPPED",
		//	          "DELETE_FAILED",
		//	          "STOP_FAILED",
		//	          "NONE"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "StatusDetails": {
		//	        "additionalProperties": false,
		//	        "patternProperties": {
		//	          "": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "StatusMessage": {
		//	        "type": "string"
		//	      },
		//	      "StatusType": {
		//	        "enum": [
		//	          "Deployment",
		//	          "AsyncExecutions"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "LastUpdatedAt",
		//	      "StatusType"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"status_summaries": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: LastUpdatedAt
					"last_updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Status
					"status": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: StatusDetails
					"status_details":    // Pattern: ""
					schema.MapAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: StatusMessage
					"status_message": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: StatusType
					"status_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "pattern": "^[A-Za-z0-9+=@_\\/-:]+$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SSMQuickSetup::ConfigurationManager",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMQuickSetup::ConfigurationManager").WithTerraformTypeName("awscc_ssmquicksetup_configuration_manager")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"configuration_definitions": "ConfigurationDefinitions",
		"created_at":                "CreatedAt",
		"description":               "Description",
		"id":                        "id",
		"last_modified_at":          "LastModifiedAt",
		"last_updated_at":           "LastUpdatedAt",
		"local_deployment_administration_role_arn": "LocalDeploymentAdministrationRoleArn",
		"local_deployment_execution_role_name":     "LocalDeploymentExecutionRoleName",
		"manager_arn":                              "ManagerArn",
		"name":                                     "Name",
		"parameters":                               "Parameters",
		"status":                                   "Status",
		"status_details":                           "StatusDetails",
		"status_message":                           "StatusMessage",
		"status_summaries":                         "StatusSummaries",
		"status_type":                              "StatusType",
		"tags":                                     "Tags",
		"type":                                     "Type",
		"type_version":                             "TypeVersion",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}