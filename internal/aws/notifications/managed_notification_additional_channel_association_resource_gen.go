// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package notifications

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
	registry.AddResourceFactory("awscc_notifications_managed_notification_additional_channel_association", managedNotificationAdditionalChannelAssociationResource)
}

// managedNotificationAdditionalChannelAssociationResource returns the Terraform awscc_notifications_managed_notification_additional_channel_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::Notifications::ManagedNotificationAdditionalChannelAssociation resource.
func managedNotificationAdditionalChannelAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ChannelArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN identifier of the channel.\nExample: arn:aws:chatbot::123456789012:chat-configuration/slack-channel/security-ops",
		//	  "pattern": "^arn:aws:(chatbot|consoleapp|notifications-contacts):[a-zA-Z0-9-]*:[0-9]{12}:[a-zA-Z0-9-_.@]+/[a-zA-Z0-9/_.@:-]+$",
		//	  "type": "string"
		//	}
		"channel_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN identifier of the channel.\nExample: arn:aws:chatbot::123456789012:chat-configuration/slack-channel/security-ops",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws:(chatbot|consoleapp|notifications-contacts):[a-zA-Z0-9-]*:[0-9]{12}:[a-zA-Z0-9-_.@]+/[a-zA-Z0-9/_.@:-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ManagedNotificationConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN identifier of the Managed Notification.\nExample: arn:aws:notifications::381491923782:managed-notification-configuration/category/AWS-Health/sub-category/Billing",
		//	  "pattern": "^arn:[-.a-z0-9]{1,63}:notifications::[0-9]{12}:managed-notification-configuration/category/[a-zA-Z0-9-]{3,64}/sub-category/[a-zA-Z0-9-]{3,64}$",
		//	  "type": "string"
		//	}
		"managed_notification_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN identifier of the Managed Notification.\nExample: arn:aws:notifications::381491923782:managed-notification-configuration/category/AWS-Health/sub-category/Billing",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:[-.a-z0-9]{1,63}:notifications::[0-9]{12}:managed-notification-configuration/category/[a-zA-Z0-9-]{3,64}/sub-category/[a-zA-Z0-9-]{3,64}$"), ""),
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
		Description: "Resource Type definition for AWS::Notifications::ManagedNotificationAdditionalChannelAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Notifications::ManagedNotificationAdditionalChannelAssociation").WithTerraformTypeName("awscc_notifications_managed_notification_additional_channel_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"channel_arn":                            "ChannelArn",
		"managed_notification_configuration_arn": "ManagedNotificationConfigurationArn",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
