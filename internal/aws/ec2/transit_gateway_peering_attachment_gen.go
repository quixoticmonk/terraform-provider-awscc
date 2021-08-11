// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

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
	registry.AddResourceTypeFactory("aws_ec2_transit_gateway_peering_attachment", transitGatewayPeeringAttachmentResourceType)
}

// transitGatewayPeeringAttachmentResourceType returns the Terraform aws_ec2_transit_gateway_peering_attachment resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::TransitGatewayPeeringAttachment resource type.
func transitGatewayPeeringAttachmentResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The time the transit gateway peering attachment was created.",
			     "format": "date-time",
			     "type": "string"
			   }
			*/
			Description: "The time the transit gateway peering attachment was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"peer_account_id": {
			// Property: PeerAccountId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the peer account",
			     "type": "string"
			   }
			*/
			Description: "The ID of the peer account",
			Type:        types.StringType,
			Required:    true,
			// PeerAccountId is a force-new attribute.
		},
		"peer_region": {
			// Property: PeerRegion
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Peer Region",
			     "type": "string"
			   }
			*/
			Description: "Peer Region",
			Type:        types.StringType,
			Required:    true,
			// PeerRegion is a force-new attribute.
		},
		"peer_transit_gateway_id": {
			// Property: PeerTransitGatewayId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the peer transit gateway.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the peer transit gateway.",
			Type:        types.StringType,
			Required:    true,
			// PeerTransitGatewayId is a force-new attribute.
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The state of the transit gateway peering attachment. Note that the initiating state has been deprecated.",
			     "type": "string"
			   }
			*/
			Description: "The state of the transit gateway peering attachment. Note that the initiating state has been deprecated.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "Code": {
			         "description": "The status code.",
			         "type": "string"
			       },
			       "Message": {
			         "description": "The status message, if applicable.",
			         "type": "string"
			       }
			     },
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"code": {
						// Property: Code
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The status code.",
						     "type": "string"
						   }
						*/
						Description: "The status code.",
						Type:        types.StringType,
						Optional:    true,
					},
					"message": {
						// Property: Message
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The status message, if applicable.",
						     "type": "string"
						   }
						*/
						Description: "The status message, if applicable.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The tags for the transit gateway peering attachment.",
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
			           "type": "string"
			         },
			         "Value": {
			           "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
			           "type": "string"
			         }
			       },
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "The tags for the transit gateway peering attachment.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
						     "type": "string"
						   }
						*/
						Description: "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
						Type:        types.StringType,
						Optional:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
						     "type": "string"
						   }
						*/
						Description: "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"transit_gateway_attachment_id": {
			// Property: TransitGatewayAttachmentId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the transit gateway peering attachment.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the transit gateway peering attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"transit_gateway_id": {
			// Property: TransitGatewayId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ID of the transit gateway.",
			     "type": "string"
			   }
			*/
			Description: "The ID of the transit gateway.",
			Type:        types.StringType,
			Required:    true,
			// TransitGatewayId is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::EC2::TransitGatewayPeeringAttachment type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayPeeringAttachment").WithTerraformTypeName("aws_ec2_transit_gateway_peering_attachment").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_ec2_transit_gateway_peering_attachment", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}