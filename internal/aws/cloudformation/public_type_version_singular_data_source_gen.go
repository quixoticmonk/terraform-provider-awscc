// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudformation_public_type_version", publicTypeVersionDataSource)
}

// publicTypeVersionDataSource returns the Terraform awscc_cloudformation_public_type_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFormation::PublicTypeVersion resource.
func publicTypeVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Number (ARN) of the extension.",
		//	  "pattern": "arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:[0-9]{12}:type/.+",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Number (ARN) of the extension.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LogDeliveryBucket
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A url to the S3 bucket where logs for the testType run will be available",
		//	  "type": "string"
		//	}
		"log_delivery_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A url to the S3 bucket where logs for the testType run will be available",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PublicTypeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
		//	  "maxLength": 1024,
		//	  "pattern": "arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/.+",
		//	  "type": "string"
		//	}
		"public_type_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PublicVersionNumber
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version number of a public third-party extension",
		//	  "maxLength": 64,
		//	  "minLength": 5,
		//	  "type": "string"
		//	}
		"public_version_number": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version number of a public third-party extension",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PublisherId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The publisher id assigned by CloudFormation for publishing in this region.",
		//	  "maxLength": 40,
		//	  "minLength": 1,
		//	  "pattern": "[0-9a-zA-Z-]{40}",
		//	  "type": "string"
		//	}
		"publisher_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The publisher id assigned by CloudFormation for publishing in this region.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The kind of extension",
		//	  "enum": [
		//	    "RESOURCE",
		//	    "MODULE",
		//	    "HOOK"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The kind of extension",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TypeName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
		//	  "pattern": "[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}",
		//	  "type": "string"
		//	}
		"type_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TypeVersionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Number (ARN) of the extension with the versionId.",
		//	  "pattern": "arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:[0-9]{12}:type/.+",
		//	  "type": "string"
		//	}
		"type_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Number (ARN) of the extension with the versionId.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFormation::PublicTypeVersion",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::PublicTypeVersion").WithTerraformTypeName("awscc_cloudformation_public_type_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"log_delivery_bucket":   "LogDeliveryBucket",
		"public_type_arn":       "PublicTypeArn",
		"public_version_number": "PublicVersionNumber",
		"publisher_id":          "PublisherId",
		"type":                  "Type",
		"type_name":             "TypeName",
		"type_version_arn":      "TypeVersionArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
