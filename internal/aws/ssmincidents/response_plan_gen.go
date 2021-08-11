// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssmincidents

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("aws_ssmincidents_response_plan", responsePlanResourceType)
}

// responsePlanResourceType returns the Terraform aws_ssmincidents_response_plan resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSMIncidents::ResponsePlan resource type.
func responsePlanResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"actions": {
			// Property: Actions
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The list of actions.",
			     "insertionOrder": true,
			     "items": {
			       "additionalProperties": false,
			       "description": "The automation configuration to launch.",
			       "properties": {
			         "SsmAutomation": {
			           "additionalProperties": false,
			           "description": "The configuration to use when starting the SSM automation document.",
			           "properties": {
			             "DocumentName": {
			               "description": "The document name to use when starting the SSM automation document.",
			               "maxLength": 128,
			               "type": "string"
			             },
			             "DocumentVersion": {
			               "description": "The version of the document to use when starting the SSM automation document.",
			               "maxLength": 128,
			               "type": "string"
			             },
			             "Parameters": {
			               "description": "The parameters to set when starting the SSM automation document.",
			               "insertionOrder": false,
			               "items": {
			                 "additionalProperties": false,
			                 "description": "A parameter to set when starting the SSM automation document.",
			                 "properties": {
			                   "Key": {
			                     "maxLength": 50,
			                     "minLength": 1,
			                     "type": "string"
			                   },
			                   "Values": {
			                     "insertionOrder": true,
			                     "items": {
			                       "description": "A value of the parameter to set when starting the SSM automation document.",
			                       "maxLength": 10000,
			                       "type": "string"
			                     },
			                     "maxItems": 10,
			                     "minItems": 1,
			                     "type": "array",
			                     "uniqueItems": true
			                   }
			                 },
			                 "required": [
			                   "Values",
			                   "Key"
			                 ],
			                 "type": "object"
			               },
			               "maxItems": 200,
			               "type": "array",
			               "uniqueItems": true
			             },
			             "RoleArn": {
			               "description": "The role ARN to use when starting the SSM automation document.",
			               "maxLength": 1000,
			               "pattern": "",
			               "type": "string"
			             },
			             "TargetAccount": {
			               "description": "The account type to use when starting the SSM automation document.",
			               "enum": [
			                 "IMPACTED_ACCOUNT",
			                 "RESPONSE_PLAN_OWNER_ACCOUNT"
			               ],
			               "type": "string"
			             }
			           },
			           "required": [
			             "RoleArn",
			             "DocumentName"
			           ],
			           "type": "object"
			         }
			       },
			       "type": "object"
			     },
			     "maxItems": 1,
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "The list of actions.",
			// Ordered set.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"ssm_automation": {
						// Property: SsmAutomation
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "The configuration to use when starting the SSM automation document.",
						     "properties": {
						       "DocumentName": {
						         "description": "The document name to use when starting the SSM automation document.",
						         "maxLength": 128,
						         "type": "string"
						       },
						       "DocumentVersion": {
						         "description": "The version of the document to use when starting the SSM automation document.",
						         "maxLength": 128,
						         "type": "string"
						       },
						       "Parameters": {
						         "description": "The parameters to set when starting the SSM automation document.",
						         "insertionOrder": false,
						         "items": {
						           "additionalProperties": false,
						           "description": "A parameter to set when starting the SSM automation document.",
						           "properties": {
						             "Key": {
						               "maxLength": 50,
						               "minLength": 1,
						               "type": "string"
						             },
						             "Values": {
						               "insertionOrder": true,
						               "items": {
						                 "description": "A value of the parameter to set when starting the SSM automation document.",
						                 "maxLength": 10000,
						                 "type": "string"
						               },
						               "maxItems": 10,
						               "minItems": 1,
						               "type": "array",
						               "uniqueItems": true
						             }
						           },
						           "required": [
						             "Values",
						             "Key"
						           ],
						           "type": "object"
						         },
						         "maxItems": 200,
						         "type": "array",
						         "uniqueItems": true
						       },
						       "RoleArn": {
						         "description": "The role ARN to use when starting the SSM automation document.",
						         "maxLength": 1000,
						         "pattern": "",
						         "type": "string"
						       },
						       "TargetAccount": {
						         "description": "The account type to use when starting the SSM automation document.",
						         "enum": [
						           "IMPACTED_ACCOUNT",
						           "RESPONSE_PLAN_OWNER_ACCOUNT"
						         ],
						         "type": "string"
						       }
						     },
						     "required": [
						       "RoleArn",
						       "DocumentName"
						     ],
						     "type": "object"
						   }
						*/
						Description: "The configuration to use when starting the SSM automation document.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"document_name": {
									// Property: DocumentName
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The document name to use when starting the SSM automation document.",
									     "maxLength": 128,
									     "type": "string"
									   }
									*/
									Description: "The document name to use when starting the SSM automation document.",
									Type:        types.StringType,
									Required:    true,
								},
								"document_version": {
									// Property: DocumentVersion
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The version of the document to use when starting the SSM automation document.",
									     "maxLength": 128,
									     "type": "string"
									   }
									*/
									Description: "The version of the document to use when starting the SSM automation document.",
									Type:        types.StringType,
									Optional:    true,
								},
								"parameters": {
									// Property: Parameters
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The parameters to set when starting the SSM automation document.",
									     "insertionOrder": false,
									     "items": {
									       "additionalProperties": false,
									       "description": "A parameter to set when starting the SSM automation document.",
									       "properties": {
									         "Key": {
									           "maxLength": 50,
									           "minLength": 1,
									           "type": "string"
									         },
									         "Values": {
									           "insertionOrder": true,
									           "items": {
									             "description": "A value of the parameter to set when starting the SSM automation document.",
									             "maxLength": 10000,
									             "type": "string"
									           },
									           "maxItems": 10,
									           "minItems": 1,
									           "type": "array",
									           "uniqueItems": true
									         }
									       },
									       "required": [
									         "Values",
									         "Key"
									       ],
									       "type": "object"
									     },
									     "maxItems": 200,
									     "type": "array",
									     "uniqueItems": true
									   }
									*/
									Description: "The parameters to set when starting the SSM automation document.",
									Attributes: providertypes.SetNestedAttributes(
										map[string]schema.Attribute{
											"key": {
												// Property: Key
												// CloudFormation resource type schema:
												/*
												   {
												     "maxLength": 50,
												     "minLength": 1,
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Required: true,
											},
											"values": {
												// Property: Values
												// CloudFormation resource type schema:
												/*
												   {
												     "insertionOrder": true,
												     "items": {
												       "description": "A value of the parameter to set when starting the SSM automation document.",
												       "maxLength": 10000,
												       "type": "string"
												     },
												     "maxItems": 10,
												     "minItems": 1,
												     "type": "array",
												     "uniqueItems": true
												   }
												*/
												// Ordered set.
												Type:     types.ListType{ElemType: types.StringType},
												Required: true,
											},
										},
										providertypes.SetNestedAttributesOptions{
											MaxItems: 200,
										},
									),
									Optional: true,
								},
								"role_arn": {
									// Property: RoleArn
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The role ARN to use when starting the SSM automation document.",
									     "maxLength": 1000,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Description: "The role ARN to use when starting the SSM automation document.",
									Type:        types.StringType,
									Required:    true,
								},
								"target_account": {
									// Property: TargetAccount
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The account type to use when starting the SSM automation document.",
									     "enum": [
									       "IMPACTED_ACCOUNT",
									       "RESPONSE_PLAN_OWNER_ACCOUNT"
									     ],
									     "type": "string"
									   }
									*/
									Description: "The account type to use when starting the SSM automation document.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 1,
				},
			),
			Optional: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the response plan.",
			     "maxLength": 1000,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The ARN of the response plan.",
			Type:        types.StringType,
			Computed:    true,
		},
		"chat_channel": {
			// Property: ChatChannel
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The chat channel configuration.",
			     "properties": {
			       "ChatbotSns": {
			         "insertionOrder": true,
			         "items": {
			           "description": "The ARN of the Chatbot SNS topic.",
			           "maxLength": 1000,
			           "pattern": "",
			           "type": "string"
			         },
			         "type": "array",
			         "uniqueItems": true
			       }
			     },
			     "type": "object"
			   }
			*/
			Description: "The chat channel configuration.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"chatbot_sns": {
						// Property: ChatbotSns
						// CloudFormation resource type schema:
						/*
						   {
						     "insertionOrder": true,
						     "items": {
						       "description": "The ARN of the Chatbot SNS topic.",
						       "maxLength": 1000,
						       "pattern": "",
						       "type": "string"
						     },
						     "type": "array",
						     "uniqueItems": true
						   }
						*/
						// Ordered set.
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"display_name": {
			// Property: DisplayName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The display name of the response plan.",
			     "maxLength": 200,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "The display name of the response plan.",
			Type:        types.StringType,
			Optional:    true,
		},
		"engagements": {
			// Property: Engagements
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The list of engagements to use.",
			     "insertionOrder": false,
			     "items": {
			       "description": "The ARN of the contact.",
			       "maxLength": 1000,
			       "pattern": "",
			       "type": "string"
			     },
			     "maxItems": 5,
			     "minItems": 1,
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "The list of engagements to use.",
			Type:        providertypes.SetType{ElemType: types.StringType},
			Optional:    true,
		},
		"incident_template": {
			// Property: IncidentTemplate
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The incident template configuration.",
			     "properties": {
			       "DedupeString": {
			         "description": "The deduplication string.",
			         "maxLength": 1000,
			         "minLength": 1,
			         "type": "string"
			       },
			       "Impact": {
			         "description": "The impact value.",
			         "type": "integer"
			       },
			       "NotificationTargets": {
			         "description": "The list of notification targets.",
			         "items": {
			           "additionalProperties": false,
			           "description": "A notification target.",
			           "properties": {
			             "SnsTopicArn": {
			               "description": "The ARN of the Chatbot SNS topic.",
			               "maxLength": 1000,
			               "pattern": "",
			               "type": "string"
			             }
			           },
			           "type": "object"
			         },
			         "maxItems": 10,
			         "type": "array"
			       },
			       "Summary": {
			         "description": "The summary string.",
			         "maxLength": 4000,
			         "minLength": 1,
			         "type": "string"
			       },
			       "Title": {
			         "description": "The title string.",
			         "maxLength": 200,
			         "type": "string"
			       }
			     },
			     "required": [
			       "Title",
			       "Impact"
			     ],
			     "type": "object"
			   }
			*/
			Description: "The incident template configuration.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"dedupe_string": {
						// Property: DedupeString
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The deduplication string.",
						     "maxLength": 1000,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The deduplication string.",
						Type:        types.StringType,
						Optional:    true,
					},
					"impact": {
						// Property: Impact
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The impact value.",
						     "type": "integer"
						   }
						*/
						Description: "The impact value.",
						Type:        types.NumberType,
						Required:    true,
					},
					"notification_targets": {
						// Property: NotificationTargets
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The list of notification targets.",
						     "items": {
						       "additionalProperties": false,
						       "description": "A notification target.",
						       "properties": {
						         "SnsTopicArn": {
						           "description": "The ARN of the Chatbot SNS topic.",
						           "maxLength": 1000,
						           "pattern": "",
						           "type": "string"
						         }
						       },
						       "type": "object"
						     },
						     "maxItems": 10,
						     "type": "array"
						   }
						*/
						Description: "The list of notification targets.",
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"sns_topic_arn": {
									// Property: SnsTopicArn
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The ARN of the Chatbot SNS topic.",
									     "maxLength": 1000,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Description: "The ARN of the Chatbot SNS topic.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
							schema.ListNestedAttributesOptions{
								MaxItems: 10,
							},
						),
						Optional: true,
					},
					"summary": {
						// Property: Summary
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The summary string.",
						     "maxLength": 4000,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The summary string.",
						Type:        types.StringType,
						Optional:    true,
					},
					"title": {
						// Property: Title
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The title string.",
						     "maxLength": 200,
						     "type": "string"
						   }
						*/
						Description: "The title string.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Required: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the response plan.",
			     "maxLength": 200,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name of the response plan.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The tags to apply to the response plan.",
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to tag a resource.",
			       "properties": {
			         "Key": {
			           "maxLength": 128,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "Value": {
			           "maxLength": 256,
			           "minLength": 1,
			           "type": "string"
			         }
			       },
			       "required": [
			         "Value",
			         "Key"
			       ],
			       "type": "object"
			     },
			     "maxItems": 50,
			     "minItems": 1,
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "The tags to apply to the response plan.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 128,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 256,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MinItems: 1,
					MaxItems: 50,
				},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource type definition for AWS::SSMIncidents::ResponsePlan",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMIncidents::ResponsePlan").WithTerraformTypeName("aws_ssmincidents_response_plan").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_ssmincidents_response_plan", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}