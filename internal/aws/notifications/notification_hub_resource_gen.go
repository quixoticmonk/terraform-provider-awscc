// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package notifications

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_notifications_notification_hub", notificationHubResource)
}

// notificationHubResource returns the Terraform awscc_notifications_notification_hub resource.
// This Terraform resource corresponds to the CloudFormation AWS::Notifications::NotificationHub resource.
func notificationHubResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NotificationHubStatusSummary
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "NotificationHubStatus": {
		//	      "enum": [
		//	        "ACTIVE",
		//	        "REGISTERING",
		//	        "DEREGISTERING",
		//	        "INACTIVE"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "NotificationHubStatusReason": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "NotificationHubStatus",
		//	    "NotificationHubStatusReason"
		//	  ],
		//	  "type": "object"
		//	}
		"notification_hub_status_summary": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: NotificationHubStatus
				"notification_hub_status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: NotificationHubStatusReason
				"notification_hub_status_reason": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Region
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Region that NotificationHub is present in.",
		//	  "maxLength": 25,
		//	  "minLength": 2,
		//	  "pattern": "^([a-z]{1,2})-([a-z]{1,15}-)+([0-9])$",
		//	  "type": "string"
		//	}
		"region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Region that NotificationHub is present in.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(2, 25),
				stringvalidator.RegexMatches(regexp.MustCompile("^([a-z]{1,2})-([a-z]{1,15}-)+([0-9])$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::Notifications::NotificationHub",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Notifications::NotificationHub").WithTerraformTypeName("awscc_notifications_notification_hub")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_time":                   "CreationTime",
		"notification_hub_status":         "NotificationHubStatus",
		"notification_hub_status_reason":  "NotificationHubStatusReason",
		"notification_hub_status_summary": "NotificationHubStatusSummary",
		"region":                          "Region",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
