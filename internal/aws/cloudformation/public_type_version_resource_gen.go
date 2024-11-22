// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_cloudformation_public_type_version", publicTypeVersionResource)
}

// publicTypeVersionResource returns the Terraform awscc_cloudformation_public_type_version resource.
// This Terraform resource corresponds to the CloudFormation AWS::CloudFormation::PublicTypeVersion resource.
func publicTypeVersionResource(ctx context.Context) (resource.Resource, error) {
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
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:[0-9]{12}:type/.+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Arn is a write-only property.
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
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
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
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
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
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(5, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
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
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
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
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"RESOURCE",
					"MODULE",
					"HOOK",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
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
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
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
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Test and Publish a resource that has been registered in the CloudFormation Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

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

	opts = opts.IsImmutableType(true)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Arn",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
