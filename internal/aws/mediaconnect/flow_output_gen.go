// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_mediaconnect_flow_output", flowOutputResourceType)
}

// flowOutputResourceType returns the Terraform aws_mediaconnect_flow_output resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::MediaConnect::FlowOutput resource type.
func flowOutputResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"cidr_allow_list": {
			// Property: CidrAllowList
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
			     "items": {
			       "type": "string"
			     },
			     "type": "array"
			   }
			*/
			Description: "The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A description of the output.",
			     "type": "string"
			   }
			*/
			Description: "A description of the output.",
			Type:        types.StringType,
			Optional:    true,
		},
		"destination": {
			// Property: Destination
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The address where you want to send the output.",
			     "type": "string"
			   }
			*/
			Description: "The address where you want to send the output.",
			Type:        types.StringType,
			Optional:    true,
		},
		"encryption": {
			// Property: Encryption
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "Information about the encryption of the flow.",
			     "properties": {
			       "Algorithm": {
			         "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
			         "enum": [
			           "aes128",
			           "aes192",
			           "aes256"
			         ],
			         "type": "string"
			       },
			       "KeyType": {
			         "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
			         "enum": [
			           "static-key"
			         ],
			         "type": "string"
			       },
			       "RoleArn": {
			         "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
			         "type": "string"
			       },
			       "SecretArn": {
			         "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
			         "type": "string"
			       }
			     },
			     "required": [
			       "Algorithm",
			       "RoleArn",
			       "SecretArn"
			     ],
			     "type": "object"
			   }
			*/
			Description: "Information about the encryption of the flow.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"algorithm": {
						// Property: Algorithm
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
						     "enum": [
						       "aes128",
						       "aes192",
						       "aes256"
						     ],
						     "type": "string"
						   }
						*/
						Description: "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
						Type:        types.StringType,
						Required:    true,
					},
					"key_type": {
						// Property: KeyType
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
						     "enum": [
						       "static-key"
						     ],
						     "type": "string"
						   }
						*/
						Description: "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
						Type:        types.StringType,
						Optional:    true,
					},
					"role_arn": {
						// Property: RoleArn
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
						     "type": "string"
						   }
						*/
						Description: "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
						Type:        types.StringType,
						Required:    true,
					},
					"secret_arn": {
						// Property: SecretArn
						// CloudFormation resource type schema:
						/*
						   {
						     "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
						     "type": "string"
						   }
						*/
						Description: " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
		},
		"flow_arn": {
			// Property: FlowArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			Type:        types.StringType,
			Required:    true,
		},
		"max_latency": {
			// Property: MaxLatency
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
			     "type": "integer"
			   }
			*/
			Description: "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the output. This value must be unique within the current flow.",
			     "type": "string"
			   }
			*/
			Description: "The name of the output. This value must be unique within the current flow.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"output_arn": {
			// Property: OutputArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the output.",
			     "type": "string"
			   }
			*/
			Description: "The ARN of the output.",
			Type:        types.StringType,
			Computed:    true,
		},
		"port": {
			// Property: Port
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The port to use when content is distributed to this output.",
			     "type": "integer"
			   }
			*/
			Description: "The port to use when content is distributed to this output.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"protocol": {
			// Property: Protocol
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The protocol that is used by the source or output.",
			     "enum": [
			       "zixi-push",
			       "rtp-fec",
			       "rtp",
			       "zixi-pull",
			       "rist"
			     ],
			     "type": "string"
			   }
			*/
			Description: "The protocol that is used by the source or output.",
			Type:        types.StringType,
			Required:    true,
		},
		"remote_id": {
			// Property: RemoteId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The remote ID for the Zixi-pull stream.",
			     "type": "string"
			   }
			*/
			Description: "The remote ID for the Zixi-pull stream.",
			Type:        types.StringType,
			Optional:    true,
		},
		"smoothing_latency": {
			// Property: SmoothingLatency
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.",
			     "type": "integer"
			   }
			*/
			Description: "The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"stream_id": {
			// Property: StreamId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
			     "type": "string"
			   }
			*/
			Description: "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
			Type:        types.StringType,
			Optional:    true,
		},
		"vpc_interface_attachment": {
			// Property: VpcInterfaceAttachment
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The settings for attaching a VPC interface to an output.",
			     "properties": {
			       "VpcInterfaceName": {
			         "description": "The name of the VPC interface to use for this output.",
			         "type": "string"
			       }
			     },
			     "type": "object"
			   }
			*/
			Description: "The settings for attaching a VPC interface to an output.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"vpc_interface_name": {
						// Property: VpcInterfaceName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The name of the VPC interface to use for this output.",
						     "type": "string"
						   }
						*/
						Description: "The name of the VPC interface to use for this output.",
						Type:        types.StringType,
						Optional:    true,
					},
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
		Description: "Resource schema for AWS::MediaConnect::FlowOutput",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::FlowOutput").WithTerraformTypeName("aws_mediaconnect_flow_output").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_mediaconnect_flow_output", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}