// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ses

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ses_mail_manager_ingress_point", mailManagerIngressPointDataSource)
}

// mailManagerIngressPointDataSource returns the Terraform awscc_ses_mail_manager_ingress_point data source.
// This Terraform data source corresponds to the CloudFormation AWS::SES::MailManagerIngressPoint resource.
func mailManagerIngressPointDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ARecord
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"a_record": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IngressPointArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"ingress_point_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IngressPointConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "SecretArn": {
		//	      "pattern": "^arn:(aws|aws-cn|aws-us-gov):secretsmanager:[a-z0-9-]+:\\d{12}:secret:[a-zA-Z0-9/_+=,.@-]+$",
		//	      "type": "string"
		//	    },
		//	    "SmtpPassword": {
		//	      "maxLength": 64,
		//	      "minLength": 8,
		//	      "pattern": "^[A-Za-z0-9!@#$%^\u0026*()_+\\-=\\[\\]{}|.,?]+$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"ingress_point_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SecretArn
				"secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SmtpPassword
				"smtp_password": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IngressPointId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"ingress_point_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IngressPointName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 63,
		//	  "minLength": 3,
		//	  "pattern": "^[A-Za-z0-9_\\-]+$",
		//	  "type": "string"
		//	}
		"ingress_point_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "PrivateNetworkConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "VpcEndpointId": {
		//	          "pattern": "^vpce-[a-zA-Z0-9]{17}$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "VpcEndpointId"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "PublicNetworkConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "IpType": {
		//	          "allOf": [
		//	            {},
		//	            {}
		//	          ]
		//	        }
		//	      },
		//	      "required": [
		//	        "IpType"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"network_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PrivateNetworkConfiguration
				"private_network_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: VpcEndpointId
						"vpc_endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PublicNetworkConfiguration
				"public_network_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: IpType
						"ip_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							CustomType: jsontypes.NormalizedType{},
							Computed:   true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RuleSetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"rule_set_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "PROVISIONING",
		//	    "DEPROVISIONING",
		//	    "UPDATING",
		//	    "ACTIVE",
		//	    "CLOSED",
		//	    "FAILED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: StatusToUpdate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ACTIVE",
		//	    "CLOSED"
		//	  ],
		//	  "type": "string"
		//	}
		"status_to_update": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9/_\\+=\\.:@\\-]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^[a-zA-Z0-9/_\\+=\\.:@\\-]*$",
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
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TrafficPolicyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"traffic_policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "OPEN",
		//	    "AUTH"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SES::MailManagerIngressPoint",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::MailManagerIngressPoint").WithTerraformTypeName("awscc_ses_mail_manager_ingress_point")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"a_record":                      "ARecord",
		"ingress_point_arn":             "IngressPointArn",
		"ingress_point_configuration":   "IngressPointConfiguration",
		"ingress_point_id":              "IngressPointId",
		"ingress_point_name":            "IngressPointName",
		"ip_type":                       "IpType",
		"key":                           "Key",
		"network_configuration":         "NetworkConfiguration",
		"private_network_configuration": "PrivateNetworkConfiguration",
		"public_network_configuration":  "PublicNetworkConfiguration",
		"rule_set_id":                   "RuleSetId",
		"secret_arn":                    "SecretArn",
		"smtp_password":                 "SmtpPassword",
		"status":                        "Status",
		"status_to_update":              "StatusToUpdate",
		"tags":                          "Tags",
		"traffic_policy_id":             "TrafficPolicyId",
		"type":                          "Type",
		"value":                         "Value",
		"vpc_endpoint_id":               "VpcEndpointId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
