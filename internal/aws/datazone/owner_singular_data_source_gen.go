// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datazone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_datazone_owner", ownerDataSource)
}

// ownerDataSource returns the Terraform awscc_datazone_owner data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataZone::Owner resource.
func ownerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DomainIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the domain in which you want to add the entity owner.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the domain in which you want to add the entity owner.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EntityIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the entity to which you want to add an owner.",
		//	  "type": "string"
		//	}
		"entity_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the entity to which you want to add an owner.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EntityType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of an entity.",
		//	  "enum": [
		//	    "DOMAIN_UNIT"
		//	  ],
		//	  "type": "string"
		//	}
		"entity_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of an entity.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Owner
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The owner that you want to add to the entity.",
		//	  "properties": {
		//	    "Group": {
		//	      "additionalProperties": false,
		//	      "description": "The properties of the domain unit owners group.",
		//	      "properties": {
		//	        "GroupIdentifier": {
		//	          "description": "The ID of the domain unit owners group.",
		//	          "pattern": "(^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$|[\\p{L}\\p{M}\\p{S}\\p{N}\\p{P}\\t\\n\\r ]+)",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "User": {
		//	      "additionalProperties": false,
		//	      "description": "The properties of the owner user.",
		//	      "properties": {
		//	        "UserIdentifier": {
		//	          "description": "The ID of the owner user.",
		//	          "pattern": "(^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$|^[a-zA-Z_0-9+=,.@-]+$|^arn:aws:iam::\\d{12}:.+$)",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"owner": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Group
				"group": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: GroupIdentifier
						"group_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ID of the domain unit owners group.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The properties of the domain unit owners group.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: User
				"user": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: UserIdentifier
						"user_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ID of the owner user.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The properties of the owner user.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The owner that you want to add to the entity.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::DataZone::Owner",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataZone::Owner").WithTerraformTypeName("awscc_datazone_owner")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"domain_identifier": "DomainIdentifier",
		"entity_identifier": "EntityIdentifier",
		"entity_type":       "EntityType",
		"group":             "Group",
		"group_identifier":  "GroupIdentifier",
		"owner":             "Owner",
		"user":              "User",
		"user_identifier":   "UserIdentifier",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
