// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ecr_pull_through_cache_rule", pullThroughCacheRuleDataSource)
}

// pullThroughCacheRuleDataSource returns the Terraform awscc_ecr_pull_through_cache_rule data source.
// This Terraform data source corresponds to the CloudFormation AWS::ECR::PullThroughCacheRule resource.
func pullThroughCacheRuleDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CredentialArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Secrets Manager secret associated with the pull through cache rule.",
		//	  "maxLength": 612,
		//	  "minLength": 50,
		//	  "pattern": "^arn:aws:secretsmanager:[a-zA-Z0-9-:]+:secret:ecr\\-pullthroughcache\\/[a-zA-Z0-9\\/_+=.@-]+$",
		//	  "type": "string"
		//	}
		"credential_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Secrets Manager secret associated with the pull through cache rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CustomRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the IAM role associated with the pull through cache rule.",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"custom_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the IAM role associated with the pull through cache rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EcrRepositoryPrefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon ECR repository prefix associated with the pull through cache rule.",
		//	  "maxLength": 30,
		//	  "minLength": 2,
		//	  "pattern": "^((?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*/?|ROOT)$",
		//	  "type": "string"
		//	}
		"ecr_repository_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon ECR repository prefix associated with the pull through cache rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpstreamRegistry
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the upstream source registry associated with the pull through cache rule.",
		//	  "type": "string"
		//	}
		"upstream_registry": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the upstream source registry associated with the pull through cache rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpstreamRegistryUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The upstream registry URL associated with the pull through cache rule.",
		//	  "type": "string"
		//	}
		"upstream_registry_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The upstream registry URL associated with the pull through cache rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpstreamRepositoryPrefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The upstream repository prefix associated with the pull through cache rule.",
		//	  "maxLength": 30,
		//	  "minLength": 2,
		//	  "pattern": "^((?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*/?|ROOT)$",
		//	  "type": "string"
		//	}
		"upstream_repository_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The upstream repository prefix associated with the pull through cache rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ECR::PullThroughCacheRule",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::PullThroughCacheRule").WithTerraformTypeName("awscc_ecr_pull_through_cache_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"credential_arn":             "CredentialArn",
		"custom_role_arn":            "CustomRoleArn",
		"ecr_repository_prefix":      "EcrRepositoryPrefix",
		"upstream_registry":          "UpstreamRegistry",
		"upstream_registry_url":      "UpstreamRegistryUrl",
		"upstream_repository_prefix": "UpstreamRepositoryPrefix",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
