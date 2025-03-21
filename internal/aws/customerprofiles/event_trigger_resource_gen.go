// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package customerprofiles

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_customerprofiles_event_trigger", eventTriggerResource)
}

// eventTriggerResource returns the Terraform awscc_customerprofiles_event_trigger resource.
// This Terraform resource corresponds to the CloudFormation AWS::CustomerProfiles::EventTrigger resource.
func eventTriggerResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when the event trigger was created.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp of when the event trigger was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the event trigger.",
		//	  "maxLength": 1000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the event trigger.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DomainName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique name of the domain.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique name of the domain.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventTriggerConditions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of conditions that determine when an event should trigger the destination.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies the circumstances under which the event should trigger the destination.",
		//	    "properties": {
		//	      "EventTriggerDimensions": {
		//	        "description": "A list of dimensions to be evaluated for the event.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "A specific event dimension to be assessed.",
		//	          "properties": {
		//	            "ObjectAttributes": {
		//	              "description": "A list of object attributes to be evaluated.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "description": "The criteria that a specific object attribute must meet to trigger the destination.",
		//	                "properties": {
		//	                  "ComparisonOperator": {
		//	                    "description": "The operator used to compare an attribute against a list of values.",
		//	                    "enum": [
		//	                      "INCLUSIVE",
		//	                      "EXCLUSIVE",
		//	                      "CONTAINS",
		//	                      "BEGINS_WITH",
		//	                      "ENDS_WITH",
		//	                      "GREATER_THAN",
		//	                      "LESS_THAN",
		//	                      "GREATER_THAN_OR_EQUAL",
		//	                      "LESS_THAN_OR_EQUAL",
		//	                      "EQUAL",
		//	                      "BEFORE",
		//	                      "AFTER",
		//	                      "ON",
		//	                      "BETWEEN",
		//	                      "NOT_BETWEEN"
		//	                    ],
		//	                    "type": "string"
		//	                  },
		//	                  "FieldName": {
		//	                    "description": "A field defined within an object type.",
		//	                    "maxLength": 64,
		//	                    "minLength": 1,
		//	                    "pattern": "^[a-zA-Z0-9_.-]+$",
		//	                    "type": "string"
		//	                  },
		//	                  "Source": {
		//	                    "description": "An attribute contained within a source object.",
		//	                    "maxLength": 1000,
		//	                    "minLength": 1,
		//	                    "type": "string"
		//	                  },
		//	                  "Values": {
		//	                    "description": "A list of attribute values used for comparison.",
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "maxLength": 255,
		//	                      "minLength": 1,
		//	                      "type": "string"
		//	                    },
		//	                    "maxItems": 10,
		//	                    "minItems": 1,
		//	                    "type": "array"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "ComparisonOperator",
		//	                  "Values"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "maxItems": 10,
		//	              "minItems": 1,
		//	              "type": "array"
		//	            }
		//	          },
		//	          "required": [
		//	            "ObjectAttributes"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "maxItems": 10,
		//	        "minItems": 1,
		//	        "type": "array"
		//	      },
		//	      "LogicalOperator": {
		//	        "description": "The operator used to combine multiple dimensions.",
		//	        "enum": [
		//	          "ANY",
		//	          "ALL",
		//	          "NONE"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "EventTriggerDimensions",
		//	      "LogicalOperator"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"event_trigger_conditions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: EventTriggerDimensions
					"event_trigger_dimensions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ObjectAttributes
								"object_attributes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: ComparisonOperator
											"comparison_operator": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The operator used to compare an attribute against a list of values.",
												Required:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.OneOf(
														"INCLUSIVE",
														"EXCLUSIVE",
														"CONTAINS",
														"BEGINS_WITH",
														"ENDS_WITH",
														"GREATER_THAN",
														"LESS_THAN",
														"GREATER_THAN_OR_EQUAL",
														"LESS_THAN_OR_EQUAL",
														"EQUAL",
														"BEFORE",
														"AFTER",
														"ON",
														"BETWEEN",
														"NOT_BETWEEN",
													),
												}, /*END VALIDATORS*/
											}, /*END ATTRIBUTE*/
											// Property: FieldName
											"field_name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "A field defined within an object type.",
												Optional:    true,
												Computed:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(1, 64),
													stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_.-]+$"), ""),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Source
											"source": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "An attribute contained within a source object.",
												Optional:    true,
												Computed:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(1, 1000),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Values
											"values": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Description: "A list of attribute values used for comparison.",
												Required:    true,
												Validators: []validator.List{ /*START VALIDATORS*/
													listvalidator.SizeBetween(1, 10),
													listvalidator.ValueStringsAre(
														stringvalidator.LengthBetween(1, 255),
													),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
													generic.Multiset(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
									}, /*END NESTED OBJECT*/
									Description: "A list of object attributes to be evaluated.",
									Required:    true,
									Validators: []validator.List{ /*START VALIDATORS*/
										listvalidator.SizeBetween(1, 10),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										generic.Multiset(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "A list of dimensions to be evaluated for the event.",
						Required:    true,
						Validators: []validator.List{ /*START VALIDATORS*/
							listvalidator.SizeBetween(1, 10),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							generic.Multiset(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LogicalOperator
					"logical_operator": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The operator used to combine multiple dimensions.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"ANY",
								"ALL",
								"NONE",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of conditions that determine when an event should trigger the destination.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 5),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventTriggerLimits
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Defines limits controlling whether an event triggers the destination, based on ingestion latency and the number of invocations per profile over specific time periods.",
		//	  "properties": {
		//	    "EventExpiration": {
		//	      "description": "Specifies that an event will only trigger the destination if it is processed within a certain latency period.",
		//	      "format": "int64",
		//	      "type": "integer"
		//	    },
		//	    "Periods": {
		//	      "description": "A list of time periods during which the limits apply.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Defines a limit and the time period during which it is enforced.",
		//	        "properties": {
		//	          "MaxInvocationsPerProfile": {
		//	            "description": "The maximum allowed number of destination invocations per profile.",
		//	            "maximum": 1000,
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          },
		//	          "Unit": {
		//	            "description": "The unit of time.",
		//	            "enum": [
		//	              "HOURS",
		//	              "DAYS",
		//	              "WEEKS",
		//	              "MONTHS"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Unlimited": {
		//	            "description": "If set to true, there is no limit on the number of destination invocations per profile. The default is false.",
		//	            "type": "boolean"
		//	          },
		//	          "Value": {
		//	            "description": "The amount of time of the specified unit.",
		//	            "maximum": 24,
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "required": [
		//	          "Unit",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 4,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"event_trigger_limits": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EventExpiration
				"event_expiration": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Specifies that an event will only trigger the destination if it is processed within a certain latency period.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Periods
				"periods": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: MaxInvocationsPerProfile
							"max_invocations_per_profile": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The maximum allowed number of destination invocations per profile.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.Between(1, 1000),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Unit
							"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The unit of time.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"HOURS",
										"DAYS",
										"WEEKS",
										"MONTHS",
									),
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Unlimited
							"unlimited": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "If set to true, there is no limit on the number of destination invocations per profile. The default is false.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
									boolplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The amount of time of the specified unit.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.Between(1, 24),
									fwvalidators.NotNullInt64(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "A list of time periods during which the limits apply.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(1, 4),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Defines limits controlling whether an event triggers the destination, based on ingestion latency and the number of invocations per profile over specific time periods.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventTriggerName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique name of the event trigger.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"event_trigger_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique name of the event trigger.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when the event trigger was most recently updated.",
		//	  "type": "string"
		//	}
		"last_updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp of when the event trigger was most recently updated.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ObjectTypeName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique name of the object type.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z_][a-zA-Z_0-9-]*$",
		//	  "type": "string"
		//	}
		"object_type_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique name of the object type.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z_][a-zA-Z_0-9-]*$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: SegmentFilter
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The destination is triggered only for profiles that meet the criteria of a segment definition.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"segment_filter": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The destination is triggered only for profiles that meet the criteria of a segment definition.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
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
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(0, 50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
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
		Description: "An event trigger resource of Amazon Connect Customer Profiles",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CustomerProfiles::EventTrigger").WithTerraformTypeName("awscc_customerprofiles_event_trigger")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"comparison_operator":         "ComparisonOperator",
		"created_at":                  "CreatedAt",
		"description":                 "Description",
		"domain_name":                 "DomainName",
		"event_expiration":            "EventExpiration",
		"event_trigger_conditions":    "EventTriggerConditions",
		"event_trigger_dimensions":    "EventTriggerDimensions",
		"event_trigger_limits":        "EventTriggerLimits",
		"event_trigger_name":          "EventTriggerName",
		"field_name":                  "FieldName",
		"key":                         "Key",
		"last_updated_at":             "LastUpdatedAt",
		"logical_operator":            "LogicalOperator",
		"max_invocations_per_profile": "MaxInvocationsPerProfile",
		"object_attributes":           "ObjectAttributes",
		"object_type_name":            "ObjectTypeName",
		"periods":                     "Periods",
		"segment_filter":              "SegmentFilter",
		"source":                      "Source",
		"tags":                        "Tags",
		"unit":                        "Unit",
		"unlimited":                   "Unlimited",
		"value":                       "Value",
		"values":                      "Values",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
