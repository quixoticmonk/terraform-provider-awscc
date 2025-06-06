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
	registry.AddResourceFactory("awscc_notifications_managed_notification_account_contact_association", managedNotificationAccountContactAssociationResource)
}

// managedNotificationAccountContactAssociationResource returns the Terraform awscc_notifications_managed_notification_account_contact_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::Notifications::ManagedNotificationAccountContactAssociation resource.
func managedNotificationAccountContactAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ContactIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This unique identifier for Contact",
		//	  "enum": [
		//	    "ACCOUNT_PRIMARY",
		//	    "ACCOUNT_ALTERNATE_SECURITY",
		//	    "ACCOUNT_ALTERNATE_OPERATIONS",
		//	    "ACCOUNT_ALTERNATE_BILLING"
		//	  ],
		//	  "type": "string"
		//	}
		"contact_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This unique identifier for Contact",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ACCOUNT_PRIMARY",
					"ACCOUNT_ALTERNATE_SECURITY",
					"ACCOUNT_ALTERNATE_OPERATIONS",
					"ACCOUNT_ALTERNATE_BILLING",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ManagedNotificationConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The managed notification configuration ARN, against which the account contact association will be created",
		//	  "pattern": "^arn:[-.a-z0-9]{1,63}:notifications::[0-9]{12}:managed-notification-configuration/category/[a-zA-Z0-9-]{3,64}/sub-category/[a-zA-Z0-9-]{3,64}$",
		//	  "type": "string"
		//	}
		"managed_notification_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The managed notification configuration ARN, against which the account contact association will be created",
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
		Description: "Resource Type definition for ManagedNotificationAccountContactAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Notifications::ManagedNotificationAccountContactAssociation").WithTerraformTypeName("awscc_notifications_managed_notification_account_contact_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"contact_identifier":                     "ContactIdentifier",
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
