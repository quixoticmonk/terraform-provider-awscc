// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package m2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_m2_deployment", deploymentDataSource)
}

// deploymentDataSource returns the Terraform awscc_m2_deployment data source.
// This Terraform data source corresponds to the CloudFormation AWS::M2::Deployment resource.
func deploymentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The application ID.",
		//	  "pattern": "^\\S{1,80}$",
		//	  "type": "string"
		//	}
		"application_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The application ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ApplicationVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version number of the application to deploy",
		//	  "type": "integer"
		//	}
		"application_version": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The version number of the application to deploy",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeploymentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The deployment ID.",
		//	  "pattern": "^\\S{1,80}$",
		//	  "type": "string"
		//	}
		"deployment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The deployment ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The environment ID.",
		//	  "pattern": "^\\S{1,80}$",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The environment ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the deployment.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the deployment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::M2::Deployment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::M2::Deployment").WithTerraformTypeName("awscc_m2_deployment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_id":      "ApplicationId",
		"application_version": "ApplicationVersion",
		"deployment_id":       "DeploymentId",
		"environment_id":      "EnvironmentId",
		"status":              "Status",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
