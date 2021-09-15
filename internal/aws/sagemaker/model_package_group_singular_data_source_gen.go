// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_sagemaker_model_package_group", modelPackageGroupDataSourceType)
}

// modelPackageGroupDataSourceType returns the Terraform awscc_sagemaker_model_package_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SageMaker::ModelPackageGroup resource type.
func modelPackageGroupDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time at which the model package group was created.",
			//   "type": "string"
			// }
			Description: "The time at which the model package group was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"model_package_group_arn": {
			// Property: ModelPackageGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the model package group.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the model package group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"model_package_group_description": {
			// Property: ModelPackageGroupDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the model package group.",
			//   "maxLength": 1024,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The description of the model package group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"model_package_group_name": {
			// Property: ModelPackageGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the model package group.",
			//   "maxLength": 63,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the model package group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"model_package_group_policy": {
			// Property: ModelPackageGroupPolicy
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"model_package_group_status": {
			// Property: ModelPackageGroupStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of a modelpackage group job.",
			//   "enum": [
			//     "Pending",
			//     "InProgress",
			//     "Completed",
			//     "Failed",
			//     "Deleting",
			//     "DeleteFailed"
			//   ],
			//   "type": "string"
			// }
			Description: "The status of a modelpackage group job.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::SageMaker::ModelPackageGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::ModelPackageGroup").WithTerraformTypeName("awscc_sagemaker_model_package_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_time":                   "CreationTime",
		"key":                             "Key",
		"model_package_group_arn":         "ModelPackageGroupArn",
		"model_package_group_description": "ModelPackageGroupDescription",
		"model_package_group_name":        "ModelPackageGroupName",
		"model_package_group_policy":      "ModelPackageGroupPolicy",
		"model_package_group_status":      "ModelPackageGroupStatus",
		"tags":                            "Tags",
		"value":                           "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}