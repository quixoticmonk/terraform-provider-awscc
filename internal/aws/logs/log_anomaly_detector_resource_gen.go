// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_logs_log_anomaly_detector", logAnomalyDetectorResource)
}

// logAnomalyDetectorResource returns the Terraform awscc_logs_log_anomaly_detector resource.
// This Terraform resource corresponds to the CloudFormation AWS::Logs::LogAnomalyDetector resource.
func logAnomalyDetectorResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Account ID for owner of detector",
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Account ID for owner of detector",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// AccountId is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: AnomalyDetectorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of LogAnomalyDetector",
		//	  "type": "string"
		//	}
		"anomaly_detector_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of LogAnomalyDetector",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AnomalyDetectorStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Current status of detector.",
		//	  "type": "string"
		//	}
		"anomaly_detector_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Current status of detector.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AnomalyVisibilityTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "number"
		//	}
		"anomaly_visibility_time": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationTimeStamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When detector was created.",
		//	  "type": "number"
		//	}
		"creation_time_stamp": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "When detector was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DetectorName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of detector",
		//	  "type": "string"
		//	}
		"detector_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of detector",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EvaluationFrequency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "How often log group is evaluated",
		//	  "enum": [
		//	    "FIVE_MIN",
		//	    "TEN_MIN",
		//	    "FIFTEEN_MIN",
		//	    "THIRTY_MIN",
		//	    "ONE_HOUR"
		//	  ],
		//	  "type": "string"
		//	}
		"evaluation_frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "How often log group is evaluated",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"FIVE_MIN",
					"TEN_MIN",
					"FIFTEEN_MIN",
					"THIRTY_MIN",
					"ONE_HOUR",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FilterPattern
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"filter_pattern": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTimeStamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When detector was lsat modified.",
		//	  "type": "number"
		//	}
		"last_modified_time_stamp": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "When detector was lsat modified.",
			Computed:    true,
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LogGroupArnList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of Arns for the given log group",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 2048,
		//	    "minLength": 20,
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"log_group_arn_list": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of Arns for the given log group",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(20, 2048),
				),
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
		Description: "The AWS::Logs::LogAnomalyDetector resource specifies a CloudWatch Logs LogAnomalyDetector.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::LogAnomalyDetector").WithTerraformTypeName("awscc_logs_log_anomaly_detector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":               "AccountId",
		"anomaly_detector_arn":     "AnomalyDetectorArn",
		"anomaly_detector_status":  "AnomalyDetectorStatus",
		"anomaly_visibility_time":  "AnomalyVisibilityTime",
		"creation_time_stamp":      "CreationTimeStamp",
		"detector_name":            "DetectorName",
		"evaluation_frequency":     "EvaluationFrequency",
		"filter_pattern":           "FilterPattern",
		"kms_key_id":               "KmsKeyId",
		"last_modified_time_stamp": "LastModifiedTimeStamp",
		"log_group_arn_list":       "LogGroupArnList",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/AccountId",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
