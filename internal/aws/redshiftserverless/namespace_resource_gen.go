// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_redshiftserverless_namespace", namespaceResource)
}

// namespaceResource returns the Terraform awscc_redshiftserverless_namespace resource.
// This Terraform resource corresponds to the CloudFormation AWS::RedshiftServerless::Namespace resource.
func namespaceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdminPasswordSecretKmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the AWS Key Management Service (KMS) key used to encrypt and store the namespace's admin credentials secret. You can only use this parameter if manageAdminPassword is true.",
		//	  "type": "string"
		//	}
		"admin_password_secret_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the AWS Key Management Service (KMS) key used to encrypt and store the namespace's admin credentials secret. You can only use this parameter if manageAdminPassword is true.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AdminUserPassword
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The password associated with the admin user for the namespace that is being created. Password must be at least 8 characters in length, should be any printable ASCII character. Must contain at least one lowercase letter, one uppercase letter and one decimal digit. You can't use adminUserPassword if manageAdminPassword is true.",
		//	  "maxLength": 64,
		//	  "minLength": 8,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"admin_user_password": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The password associated with the admin user for the namespace that is being created. Password must be at least 8 characters in length, should be any printable ASCII character. Must contain at least one lowercase letter, one uppercase letter and one decimal digit. You can't use adminUserPassword if manageAdminPassword is true.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(8, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// AdminUserPassword is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: AdminUsername
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user name associated with the admin user for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
		//	  "pattern": "[a-zA-Z][a-zA-Z_0-9+.@-]*",
		//	  "type": "string"
		//	}
		"admin_username": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user name associated with the admin user for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z][a-zA-Z_0-9+.@-]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DbName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The database name associated for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
		//	  "maxLength": 127,
		//	  "pattern": "[a-zA-Z][a-zA-Z_0-9+.@-]*",
		//	  "type": "string"
		//	}
		"db_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The database name associated for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(127),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z][a-zA-Z_0-9+.@-]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultIamRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default IAM role ARN for the namespace that is being created.",
		//	  "type": "string"
		//	}
		"default_iam_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The default IAM role ARN for the namespace that is being created.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FinalSnapshotName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the namespace the source snapshot was created from. Please specify the name if needed before deleting namespace",
		//	  "maxLength": 255,
		//	  "pattern": "[a-z][a-z0-9]*(-[a-z0-9]+)*",
		//	  "type": "string"
		//	}
		"final_snapshot_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the namespace the source snapshot was created from. Please specify the name if needed before deleting namespace",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(255),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-z][a-z0-9]*(-[a-z0-9]+)*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// FinalSnapshotName is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: FinalSnapshotRetentionPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of days to retain automated snapshot in the destination region after they are copied from the source region. If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3,653.",
		//	  "type": "integer"
		//	}
		"final_snapshot_retention_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of days to retain automated snapshot in the destination region after they are copied from the source region. If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3,653.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// FinalSnapshotRetentionPeriod is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: IamRoles
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of AWS Identity and Access Management (IAM) roles that can be used by the namespace to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. The Default role limit for each request is 10.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 512,
		//	    "minLength": 0,
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"iam_roles": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of AWS Identity and Access Management (IAM) roles that can be used by the namespace to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. The Default role limit for each request is 10.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(0, 512),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the namespace.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the namespace.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LogExports
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The collection of log types to be exported provided by the customer. Should only be one of the three supported log types: userlog, useractivitylog and connectionlog",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "enum": [
		//	      "useractivitylog",
		//	      "userlog",
		//	      "connectionlog"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "maxItems": 16,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"log_exports": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The collection of log types to be exported provided by the customer. Should only be one of the three supported log types: userlog, useractivitylog and connectionlog",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 16),
				listvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"useractivitylog",
						"userlog",
						"connectionlog",
					),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ManageAdminPassword
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If true, Amazon Redshift uses AWS Secrets Manager to manage the namespace's admin credentials. You can't use adminUserPassword if manageAdminPassword is true. If manageAdminPassword is false or not set, Amazon Redshift uses adminUserPassword for the admin user account's password.",
		//	  "type": "boolean"
		//	}
		"manage_admin_password": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "If true, Amazon Redshift uses AWS Secrets Manager to manage the namespace's admin credentials. You can't use adminUserPassword if manageAdminPassword is true. If manageAdminPassword is false or not set, Amazon Redshift uses adminUserPassword for the admin user account's password.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// ManageAdminPassword is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Namespace
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Definition of Namespace resource.",
		//	  "properties": {
		//	    "AdminPasswordSecretArn": {
		//	      "type": "string"
		//	    },
		//	    "AdminPasswordSecretKmsKeyId": {
		//	      "type": "string"
		//	    },
		//	    "AdminUsername": {
		//	      "type": "string"
		//	    },
		//	    "CreationDate": {
		//	      "type": "string"
		//	    },
		//	    "DbName": {
		//	      "pattern": "[a-zA-Z][a-zA-Z_0-9+.@-]*",
		//	      "type": "string"
		//	    },
		//	    "DefaultIamRoleArn": {
		//	      "type": "string"
		//	    },
		//	    "IamRoles": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 512,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "KmsKeyId": {
		//	      "type": "string"
		//	    },
		//	    "LogExports": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "enum": [
		//	          "useractivitylog",
		//	          "userlog",
		//	          "connectionlog"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "maxItems": 16,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "NamespaceArn": {
		//	      "type": "string"
		//	    },
		//	    "NamespaceId": {
		//	      "type": "string"
		//	    },
		//	    "NamespaceName": {
		//	      "maxLength": 64,
		//	      "minLength": 3,
		//	      "pattern": "^[a-z0-9-]+$",
		//	      "type": "string"
		//	    },
		//	    "Status": {
		//	      "enum": [
		//	        "AVAILABLE",
		//	        "MODIFYING",
		//	        "DELETING"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"namespace": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AdminPasswordSecretArn
				"admin_password_secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AdminPasswordSecretKmsKeyId
				"admin_password_secret_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AdminUsername
				"admin_username": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: CreationDate
				"creation_date": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DbName
				"db_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DefaultIamRoleArn
				"default_iam_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: IamRoles
				"iam_roles": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: KmsKeyId
				"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: LogExports
				"log_exports": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: NamespaceArn
				"namespace_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: NamespaceId
				"namespace_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: NamespaceName
				"namespace_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Definition of Namespace resource.",
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NamespaceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for the namespace. You use this identifier to refer to the namespace for any subsequent namespace operations such as deleting or modifying. All alphabetical characters must be lower case. Namespace name should be unique for all namespaces within an AWS account.",
		//	  "maxLength": 64,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z0-9-]+$",
		//	  "type": "string"
		//	}
		"namespace_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for the namespace. You use this identifier to refer to the namespace for any subsequent namespace operations such as deleting or modifying. All alphabetical characters must be lower case. Namespace name should be unique for all namespaces within an AWS account.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-z0-9-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NamespaceResourcePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The resource policy document that will be attached to the namespace.",
		//	  "type": "object"
		//	}
		"namespace_resource_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "The resource policy document that will be attached to the namespace.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RedshiftIdcApplicationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN for the Redshift application that integrates with IAM Identity Center.",
		//	  "type": "string"
		//	}
		"redshift_idc_application_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN for the Redshift application that integrates with IAM Identity Center.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// RedshiftIdcApplicationArn is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: SnapshotCopyConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The snapshot copy configurations for the namespace.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "DestinationKmsKeyId": {
		//	        "type": "string"
		//	      },
		//	      "DestinationRegion": {
		//	        "type": "string"
		//	      },
		//	      "SnapshotRetentionPeriod": {
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "required": [
		//	      "DestinationRegion"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"snapshot_copy_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DestinationKmsKeyId
					"destination_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: DestinationRegion
					"destination_region": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: SnapshotRetentionPeriod
					"snapshot_retention_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The snapshot copy configurations for the namespace.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of tags for the namespace.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
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
						Optional: true,
						Computed: true,
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
			Description: "The list of tags for the namespace.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
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
		Description: "Definition of AWS::RedshiftServerless::Namespace Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RedshiftServerless::Namespace").WithTerraformTypeName("awscc_redshiftserverless_namespace")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"admin_password_secret_arn":        "AdminPasswordSecretArn",
		"admin_password_secret_kms_key_id": "AdminPasswordSecretKmsKeyId",
		"admin_user_password":              "AdminUserPassword",
		"admin_username":                   "AdminUsername",
		"creation_date":                    "CreationDate",
		"db_name":                          "DbName",
		"default_iam_role_arn":             "DefaultIamRoleArn",
		"destination_kms_key_id":           "DestinationKmsKeyId",
		"destination_region":               "DestinationRegion",
		"final_snapshot_name":              "FinalSnapshotName",
		"final_snapshot_retention_period":  "FinalSnapshotRetentionPeriod",
		"iam_roles":                        "IamRoles",
		"key":                              "Key",
		"kms_key_id":                       "KmsKeyId",
		"log_exports":                      "LogExports",
		"manage_admin_password":            "ManageAdminPassword",
		"namespace":                        "Namespace",
		"namespace_arn":                    "NamespaceArn",
		"namespace_id":                     "NamespaceId",
		"namespace_name":                   "NamespaceName",
		"namespace_resource_policy":        "NamespaceResourcePolicy",
		"redshift_idc_application_arn":     "RedshiftIdcApplicationArn",
		"snapshot_copy_configurations":     "SnapshotCopyConfigurations",
		"snapshot_retention_period":        "SnapshotRetentionPeriod",
		"status":                           "Status",
		"tags":                             "Tags",
		"value":                            "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/AdminUserPassword",
		"/properties/FinalSnapshotName",
		"/properties/FinalSnapshotRetentionPeriod",
		"/properties/ManageAdminPassword",
		"/properties/RedshiftIdcApplicationArn",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
