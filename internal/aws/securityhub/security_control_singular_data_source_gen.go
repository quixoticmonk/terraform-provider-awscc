// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_securityhub_security_control", securityControlDataSource)
}

// securityControlDataSource returns the Terraform awscc_securityhub_security_control data source.
// This Terraform data source corresponds to the CloudFormation AWS::SecurityHub::SecurityControl resource.
func securityControlDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: LastUpdateReason
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The most recent reason for updating the customizable properties of a security control. This differs from the UpdateReason field of the BatchUpdateStandardsControlAssociations API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores.",
		//	  "pattern": "^([^\u0000-]|[-_ a-zA-Z0-9])+$",
		//	  "type": "string"
		//	}
		"last_update_reason": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The most recent reason for updating the customizable properties of a security control. This differs from the UpdateReason field of the BatchUpdateStandardsControlAssociations API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Parameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that identifies the name of a control parameter, its current value, and whether it has been customized.",
		//	  "patternProperties": {
		//	    "": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Value": {
		//	          "additionalProperties": false,
		//	          "oneOf": [
		//	            {
		//	              "required": [
		//	                "Boolean"
		//	              ]
		//	            },
		//	            {
		//	              "required": [
		//	                "Double"
		//	              ]
		//	            },
		//	            {
		//	              "required": [
		//	                "Enum"
		//	              ]
		//	            },
		//	            {
		//	              "required": [
		//	                "EnumList"
		//	              ]
		//	            },
		//	            {
		//	              "required": [
		//	                "Integer"
		//	              ]
		//	            },
		//	            {
		//	              "required": [
		//	                "IntegerList"
		//	              ]
		//	            },
		//	            {
		//	              "required": [
		//	                "String"
		//	              ]
		//	            },
		//	            {
		//	              "required": [
		//	                "StringList"
		//	              ]
		//	            }
		//	          ],
		//	          "properties": {
		//	            "Boolean": {
		//	              "description": "A control parameter that is a boolean.",
		//	              "type": "boolean"
		//	            },
		//	            "Double": {
		//	              "description": "A control parameter that is a double.",
		//	              "type": "number"
		//	            },
		//	            "Enum": {
		//	              "description": "A control parameter that is a enum.",
		//	              "pattern": ".*\\S.*",
		//	              "type": "string"
		//	            },
		//	            "EnumList": {
		//	              "description": "A control parameter that is a list of enums.",
		//	              "items": {
		//	                "pattern": ".*\\S.*",
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "Integer": {
		//	              "description": "A control parameter that is a integer.",
		//	              "type": "integer"
		//	            },
		//	            "IntegerList": {
		//	              "description": "A control parameter that is a list of integers.",
		//	              "items": {
		//	                "type": "integer"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "String": {
		//	              "description": "A control parameter that is a string.",
		//	              "pattern": ".*\\S.*",
		//	              "type": "string"
		//	            },
		//	            "StringList": {
		//	              "description": "A control parameter that is a list of strings.",
		//	              "items": {
		//	                "pattern": ".*\\S.*",
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "ValueType": {
		//	          "enum": [
		//	            "DEFAULT",
		//	            "CUSTOM"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ValueType"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"parameters":              // Pattern: ""
		schema.MapNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Value
					"value": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Boolean
							"boolean": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "A control parameter that is a boolean.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Double
							"double": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "A control parameter that is a double.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Enum
							"enum": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A control parameter that is a enum.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: EnumList
							"enum_list": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "A control parameter that is a list of enums.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Integer
							"integer": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "A control parameter that is a integer.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: IntegerList
							"integer_list": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.Int64Type,
								Description: "A control parameter that is a list of integers.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: String
							"string": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A control parameter that is a string.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: StringList
							"string_list": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "A control parameter that is a list of strings.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ValueType
					"value_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An object that identifies the name of a control parameter, its current value, and whether it has been customized.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityControlArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for a security control across standards, such as `arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1`. This parameter doesn't mention a specific standard.",
		//	  "pattern": ".*\\S.*",
		//	  "type": "string"
		//	}
		"security_control_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for a security control across standards, such as `arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1`. This parameter doesn't mention a specific standard.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityControlId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier of a security control across standards. Values for this field typically consist of an AWS service name and a number, such as APIGateway.3.",
		//	  "pattern": ".*\\S.*",
		//	  "type": "string"
		//	}
		"security_control_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier of a security control across standards. Values for this field typically consist of an AWS service name and a number, such as APIGateway.3.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SecurityHub::SecurityControl",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecurityHub::SecurityControl").WithTerraformTypeName("awscc_securityhub_security_control")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"boolean":              "Boolean",
		"double":               "Double",
		"enum":                 "Enum",
		"enum_list":            "EnumList",
		"integer":              "Integer",
		"integer_list":         "IntegerList",
		"last_update_reason":   "LastUpdateReason",
		"parameters":           "Parameters",
		"security_control_arn": "SecurityControlArn",
		"security_control_id":  "SecurityControlId",
		"string":               "String",
		"string_list":          "StringList",
		"value":                "Value",
		"value_type":           "ValueType",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
