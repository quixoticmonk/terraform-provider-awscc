// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package paymentcryptography

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_paymentcryptography_key", keyResource)
}

// keyResource returns the Terraform awscc_paymentcryptography_key resource.
// This Terraform resource corresponds to the CloudFormation AWS::PaymentCryptography::Key resource.
func keyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Enabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Exportable
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"exportable": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: KeyAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "KeyAlgorithm": {
		//	      "enum": [
		//	        "TDES_2KEY",
		//	        "TDES_3KEY",
		//	        "AES_128",
		//	        "AES_192",
		//	        "AES_256",
		//	        "HMAC_SHA256",
		//	        "HMAC_SHA384",
		//	        "HMAC_SHA512",
		//	        "HMAC_SHA224",
		//	        "RSA_2048",
		//	        "RSA_3072",
		//	        "RSA_4096",
		//	        "ECC_NIST_P256",
		//	        "ECC_NIST_P384",
		//	        "ECC_NIST_P521"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "KeyClass": {
		//	      "enum": [
		//	        "SYMMETRIC_KEY",
		//	        "ASYMMETRIC_KEY_PAIR",
		//	        "PRIVATE_KEY",
		//	        "PUBLIC_KEY"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "KeyModesOfUse": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Decrypt": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "DeriveKey": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "Encrypt": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "Generate": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "NoRestrictions": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "Sign": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "Unwrap": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "Verify": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "Wrap": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "KeyUsage": {
		//	      "enum": [
		//	        "TR31_B0_BASE_DERIVATION_KEY",
		//	        "TR31_C0_CARD_VERIFICATION_KEY",
		//	        "TR31_D0_SYMMETRIC_DATA_ENCRYPTION_KEY",
		//	        "TR31_D1_ASYMMETRIC_KEY_FOR_DATA_ENCRYPTION",
		//	        "TR31_E0_EMV_MKEY_APP_CRYPTOGRAMS",
		//	        "TR31_E1_EMV_MKEY_CONFIDENTIALITY",
		//	        "TR31_E2_EMV_MKEY_INTEGRITY",
		//	        "TR31_E4_EMV_MKEY_DYNAMIC_NUMBERS",
		//	        "TR31_E5_EMV_MKEY_CARD_PERSONALIZATION",
		//	        "TR31_E6_EMV_MKEY_OTHER",
		//	        "TR31_K0_KEY_ENCRYPTION_KEY",
		//	        "TR31_K1_KEY_BLOCK_PROTECTION_KEY",
		//	        "TR31_K3_ASYMMETRIC_KEY_FOR_KEY_AGREEMENT",
		//	        "TR31_M3_ISO_9797_3_MAC_KEY",
		//	        "TR31_M1_ISO_9797_1_MAC_KEY",
		//	        "TR31_M6_ISO_9797_5_CMAC_KEY",
		//	        "TR31_M7_HMAC_KEY",
		//	        "TR31_P0_PIN_ENCRYPTION_KEY",
		//	        "TR31_P1_PIN_GENERATION_KEY",
		//	        "TR31_S0_ASYMMETRIC_KEY_FOR_DIGITAL_SIGNATURE",
		//	        "TR31_V1_IBM3624_PIN_VERIFICATION_KEY",
		//	        "TR31_V2_VISA_PIN_VERIFICATION_KEY",
		//	        "TR31_K2_TR34_ASYMMETRIC_KEY"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "KeyAlgorithm",
		//	    "KeyClass",
		//	    "KeyModesOfUse",
		//	    "KeyUsage"
		//	  ],
		//	  "type": "object"
		//	}
		"key_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KeyAlgorithm
				"key_algorithm": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"TDES_2KEY",
							"TDES_3KEY",
							"AES_128",
							"AES_192",
							"AES_256",
							"HMAC_SHA256",
							"HMAC_SHA384",
							"HMAC_SHA512",
							"HMAC_SHA224",
							"RSA_2048",
							"RSA_3072",
							"RSA_4096",
							"ECC_NIST_P256",
							"ECC_NIST_P384",
							"ECC_NIST_P521",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: KeyClass
				"key_class": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"SYMMETRIC_KEY",
							"ASYMMETRIC_KEY_PAIR",
							"PRIVATE_KEY",
							"PUBLIC_KEY",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: KeyModesOfUse
				"key_modes_of_use": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Decrypt
						"decrypt": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: DeriveKey
						"derive_key": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Encrypt
						"encrypt": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Generate
						"generate": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: NoRestrictions
						"no_restrictions": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Sign
						"sign": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Unwrap
						"unwrap": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Verify
						"verify": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Wrap
						"wrap": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: KeyUsage
				"key_usage": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"TR31_B0_BASE_DERIVATION_KEY",
							"TR31_C0_CARD_VERIFICATION_KEY",
							"TR31_D0_SYMMETRIC_DATA_ENCRYPTION_KEY",
							"TR31_D1_ASYMMETRIC_KEY_FOR_DATA_ENCRYPTION",
							"TR31_E0_EMV_MKEY_APP_CRYPTOGRAMS",
							"TR31_E1_EMV_MKEY_CONFIDENTIALITY",
							"TR31_E2_EMV_MKEY_INTEGRITY",
							"TR31_E4_EMV_MKEY_DYNAMIC_NUMBERS",
							"TR31_E5_EMV_MKEY_CARD_PERSONALIZATION",
							"TR31_E6_EMV_MKEY_OTHER",
							"TR31_K0_KEY_ENCRYPTION_KEY",
							"TR31_K1_KEY_BLOCK_PROTECTION_KEY",
							"TR31_K3_ASYMMETRIC_KEY_FOR_KEY_AGREEMENT",
							"TR31_M3_ISO_9797_3_MAC_KEY",
							"TR31_M1_ISO_9797_1_MAC_KEY",
							"TR31_M6_ISO_9797_5_CMAC_KEY",
							"TR31_M7_HMAC_KEY",
							"TR31_P0_PIN_ENCRYPTION_KEY",
							"TR31_P1_PIN_GENERATION_KEY",
							"TR31_S0_ASYMMETRIC_KEY_FOR_DIGITAL_SIGNATURE",
							"TR31_V1_IBM3624_PIN_VERIFICATION_KEY",
							"TR31_V2_VISA_PIN_VERIFICATION_KEY",
							"TR31_K2_TR34_ASYMMETRIC_KEY",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: KeyCheckValueAlgorithm
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "CMAC",
		//	    "ANSI_X9_24"
		//	  ],
		//	  "type": "string"
		//	}
		"key_check_value_algorithm": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"CMAC",
					"ANSI_X9_24",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KeyIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 322,
		//	  "minLength": 7,
		//	  "pattern": "^arn:aws:payment-cryptography:[a-z]{2}-[a-z]{1,16}-[0-9]+:[0-9]{12}:(key/[0-9a-zA-Z]{16,64}|alias/[a-zA-Z0-9/_-]+)$|^alias/[a-zA-Z0-9/_-]+$",
		//	  "type": "string"
		//	}
		"key_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KeyOrigin
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Defines the source of a key",
		//	  "enum": [
		//	    "EXTERNAL",
		//	    "AWS_PAYMENT_CRYPTOGRAPHY"
		//	  ],
		//	  "type": "string"
		//	}
		"key_origin": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Defines the source of a key",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KeyState
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Defines the state of a key",
		//	  "enum": [
		//	    "CREATE_IN_PROGRESS",
		//	    "CREATE_COMPLETE",
		//	    "DELETE_PENDING",
		//	    "DELETE_COMPLETE"
		//	  ],
		//	  "type": "string"
		//	}
		"key_state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Defines the state of a key",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
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
		//	        "minLength": 1,
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
							stringvalidator.LengthBetween(1, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
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
		Description: "Definition of AWS::PaymentCryptography::Key Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::PaymentCryptography::Key").WithTerraformTypeName("awscc_paymentcryptography_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"decrypt":                   "Decrypt",
		"derive_key":                "DeriveKey",
		"enabled":                   "Enabled",
		"encrypt":                   "Encrypt",
		"exportable":                "Exportable",
		"generate":                  "Generate",
		"key":                       "Key",
		"key_algorithm":             "KeyAlgorithm",
		"key_attributes":            "KeyAttributes",
		"key_check_value_algorithm": "KeyCheckValueAlgorithm",
		"key_class":                 "KeyClass",
		"key_identifier":            "KeyIdentifier",
		"key_modes_of_use":          "KeyModesOfUse",
		"key_origin":                "KeyOrigin",
		"key_state":                 "KeyState",
		"key_usage":                 "KeyUsage",
		"no_restrictions":           "NoRestrictions",
		"sign":                      "Sign",
		"tags":                      "Tags",
		"unwrap":                    "Unwrap",
		"value":                     "Value",
		"verify":                    "Verify",
		"wrap":                      "Wrap",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
