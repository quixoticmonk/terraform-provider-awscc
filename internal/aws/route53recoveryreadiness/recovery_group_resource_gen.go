// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_route53recoveryreadiness_recovery_group", recoveryGroupResourceType)
}

// recoveryGroupResourceType returns the Terraform awscc_route53recoveryreadiness_recovery_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53RecoveryReadiness::RecoveryGroup resource type.
func recoveryGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cells": {
			// Property: Cells
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of the cell Amazon Resource Names (ARNs) in the recovery group.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 256,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "A list of the cell Amazon Resource Names (ARNs) in the recovery group.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 5),
			},
		},
		"recovery_group_arn": {
			// Property: RecoveryGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "A collection of tags associated with a resource.",
			Type:        types.StringType,
			Computed:    true,
		},
		"recovery_group_name": {
			// Property: RecoveryGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the recovery group to create.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the recovery group to create.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			// RecoveryGroupName is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "insertionOrder": false,
			//         "items": {
			//           "maxItems": 50,
			//           "type": "string"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A collection of tags associated with a resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "AWS Route53 Recovery Readiness Recovery Group Schema and API specifications.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryReadiness::RecoveryGroup").WithTerraformTypeName("awscc_route53recoveryreadiness_recovery_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cells":               "Cells",
		"key":                 "Key",
		"recovery_group_arn":  "RecoveryGroupArn",
		"recovery_group_name": "RecoveryGroupName",
		"tags":                "Tags",
		"value":               "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53recoveryreadiness_recovery_group", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}