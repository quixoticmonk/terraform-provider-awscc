// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package bedrock

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_bedrock_intelligent_prompt_router", intelligentPromptRouterDataSource)
}

// intelligentPromptRouterDataSource returns the Terraform awscc_bedrock_intelligent_prompt_router data source.
// This Terraform data source corresponds to the CloudFormation AWS::Bedrock::IntelligentPromptRouter resource.
func intelligentPromptRouterDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time Stamp",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Time Stamp",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of the Prompt Router.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "^([0-9a-zA-Z:.][ _-]?)+$",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of the Prompt Router.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FallbackModel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Model configuration",
		//	  "properties": {
		//	    "ModelArn": {
		//	      "description": "Arn of underlying model which are added in the Prompt Router.",
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "pattern": "(^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}::foundation-model/[a-z0-9-]{1,63}[.]{1}([a-z0-9-]{1,63}[.]){0,2}[a-z0-9-]{1,63}([:][a-z0-9-]{1,63}){0,2})|(^arn:aws(|-us-gov|-cn|-iso|-iso-b):bedrock:(|[0-9a-z-]{0,20}):(|[0-9]{12}):(inference-profile|application-inference-profile)/[a-zA-Z0-9-:.]+)$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "ModelArn"
		//	  ],
		//	  "type": "object"
		//	}
		"fallback_model": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ModelArn
				"model_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Arn of underlying model which are added in the Prompt Router.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Model configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Models
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of model configuration",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Model configuration",
		//	    "properties": {
		//	      "ModelArn": {
		//	        "description": "Arn of underlying model which are added in the Prompt Router.",
		//	        "maxLength": 2048,
		//	        "minLength": 1,
		//	        "pattern": "(^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}::foundation-model/[a-z0-9-]{1,63}[.]{1}([a-z0-9-]{1,63}[.]){0,2}[a-z0-9-]{1,63}([:][a-z0-9-]{1,63}){0,2})|(^arn:aws(|-us-gov|-cn|-iso|-iso-b):bedrock:(|[0-9a-z-]{0,20}):(|[0-9]{12}):(inference-profile|application-inference-profile)/[a-zA-Z0-9-:.]+)$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ModelArn"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"models": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ModelArn
					"model_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Arn of underlying model which are added in the Prompt Router.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of model configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PromptRouterArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn of the Prompt Router.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:(default-)?prompt-router/[a-zA-Z0-9-:.]+$",
		//	  "type": "string"
		//	}
		"prompt_router_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn of the Prompt Router.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PromptRouterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the Prompt Router.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^([0-9a-zA-Z][ _-]?)+$",
		//	  "type": "string"
		//	}
		"prompt_router_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the Prompt Router.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoutingCriteria
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Represents the criteria used for routing requests.",
		//	  "properties": {
		//	    "ResponseQualityDifference": {
		//	      "maximum": 100,
		//	      "minimum": 0,
		//	      "type": "number"
		//	    }
		//	  },
		//	  "required": [
		//	    "ResponseQualityDifference"
		//	  ],
		//	  "type": "object"
		//	}
		"routing_criteria": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ResponseQualityDifference
				"response_quality_difference": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Represents the criteria used for routing requests.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Status of a PromptRouter",
		//	  "enum": [
		//	    "AVAILABLE"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Status of a PromptRouter",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of Tags",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Definition of the key/value pair for a tag",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "Tag Key",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s._:/=+@-]*$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "Tag Value",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^[a-zA-Z0-9\\s._:/=+@-]*$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Tag Key",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Tag Value",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of Tags",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Type of a Prompt Router",
		//	  "enum": [
		//	    "custom",
		//	    "default"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Type of a Prompt Router",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time Stamp",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Time Stamp",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Bedrock::IntelligentPromptRouter",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Bedrock::IntelligentPromptRouter").WithTerraformTypeName("awscc_bedrock_intelligent_prompt_router")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_at":                  "CreatedAt",
		"description":                 "Description",
		"fallback_model":              "FallbackModel",
		"key":                         "Key",
		"model_arn":                   "ModelArn",
		"models":                      "Models",
		"prompt_router_arn":           "PromptRouterArn",
		"prompt_router_name":          "PromptRouterName",
		"response_quality_difference": "ResponseQualityDifference",
		"routing_criteria":            "RoutingCriteria",
		"status":                      "Status",
		"tags":                        "Tags",
		"type":                        "Type",
		"updated_at":                  "UpdatedAt",
		"value":                       "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
