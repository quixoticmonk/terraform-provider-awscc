// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package datazone

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
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
	registry.AddResourceFactory("awscc_datazone_domain_unit", domainUnitResource)
}

// domainUnitResource returns the Terraform awscc_datazone_domain_unit resource.
// This Terraform resource corresponds to the CloudFormation AWS::DataZone::DomainUnit resource.
func domainUnitResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp at which the domain unit was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp at which the domain unit was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the domain unit.",
		//	  "maxLength": 2048,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the domain unit.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DomainId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the domain where the domain unit was created.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the domain where the domain unit was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DomainIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the domain where you want to create a domain unit.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the domain where you want to create a domain unit.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^dzd[-_][a-zA-Z0-9_-]{1,36}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the domain unit.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9_-]+$",
		//	  "type": "string"
		//	}
		"domain_unit_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the domain unit.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Identifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the domain unit that you want to get.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9_-]+$",
		//	  "type": "string"
		//	}
		"identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the domain unit that you want to get.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp at which the domain unit was last updated.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"last_updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp at which the domain unit was last updated.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the domain unit.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[\\w -]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the domain unit.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\w -]+$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ParentDomainUnitId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the parent domain unit.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9_-]+$",
		//	  "type": "string"
		//	}
		"parent_domain_unit_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the parent domain unit.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ParentDomainUnitIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the parent domain unit.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9_-]+$",
		//	  "type": "string"
		//	}
		"parent_domain_unit_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the parent domain unit.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-z0-9_-]+$"), ""),
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
		Description: "A domain unit enables you to easily organize your assets and other domain entities under specific business units and teams.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataZone::DomainUnit").WithTerraformTypeName("awscc_datazone_domain_unit")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_at":                    "CreatedAt",
		"description":                   "Description",
		"domain_id":                     "DomainId",
		"domain_identifier":             "DomainIdentifier",
		"domain_unit_id":                "Id",
		"identifier":                    "Identifier",
		"last_updated_at":               "LastUpdatedAt",
		"name":                          "Name",
		"parent_domain_unit_id":         "ParentDomainUnitId",
		"parent_domain_unit_identifier": "ParentDomainUnitIdentifier",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
