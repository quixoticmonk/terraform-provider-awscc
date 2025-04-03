// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_mediaconnect_flow_output", flowOutputDataSource)
}

// flowOutputDataSource returns the Terraform awscc_mediaconnect_flow_output data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaConnect::FlowOutput resource.
func flowOutputDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CidrAllowList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"cidr_allow_list": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the output.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the output.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Destination
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The address where you want to send the output.",
		//	  "type": "string"
		//	}
		"destination": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The address where you want to send the output.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Encryption
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The type of key used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
		//	  "properties": {
		//	    "Algorithm": {
		//	      "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
		//	      "enum": [
		//	        "aes128",
		//	        "aes192",
		//	        "aes256"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "KeyType": {
		//	      "default": "static-key",
		//	      "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
		//	      "enum": [
		//	        "static-key",
		//	        "srt-password"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "RoleArn": {
		//	      "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
		//	      "type": "string"
		//	    },
		//	    "SecretArn": {
		//	      "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "RoleArn",
		//	    "SecretArn"
		//	  ],
		//	  "type": "object"
		//	}
		"encryption": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Algorithm
				"algorithm": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: KeyType
				"key_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RoleArn
				"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SecretArn
				"secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The type of key used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FlowArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
		//	  "type": "string"
		//	}
		"flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaxLatency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
		//	  "type": "integer"
		//	}
		"max_latency": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MediaStreamOutputConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The definition for each media stream that is associated with the output.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The media stream that is associated with the output, and the parameters for that association.",
		//	    "properties": {
		//	      "DestinationConfigurations": {
		//	        "description": "The media streams that you want to associate with the output.",
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "The definition of a media stream that is associated with the output.",
		//	          "properties": {
		//	            "DestinationIp": {
		//	              "description": "The IP address where contents of the media stream will be sent.",
		//	              "type": "string"
		//	            },
		//	            "DestinationPort": {
		//	              "description": "The port to use when the content of the media stream is distributed to the output.",
		//	              "type": "integer"
		//	            },
		//	            "Interface": {
		//	              "additionalProperties": false,
		//	              "description": "The VPC interface that is used for the media stream associated with the output.",
		//	              "properties": {
		//	                "Name": {
		//	                  "description": "The name of the VPC interface that you want to use for the media stream associated with the output.",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Name"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "required": [
		//	            "DestinationIp",
		//	            "DestinationPort",
		//	            "Interface"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "EncodingName": {
		//	        "description": "The format that will be used to encode the data. For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video streams on sources or outputs that use the CDI protocol, set the encoding name to raw. For video streams on sources or outputs that use the ST 2110 JPEG XS protocol, set the encoding name to jxsv.",
		//	        "enum": [
		//	          "jxsv",
		//	          "raw",
		//	          "smpte291",
		//	          "pcm"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "EncodingParameters": {
		//	        "additionalProperties": false,
		//	        "description": "A collection of parameters that determine how MediaConnect will convert the content. These fields only apply to outputs on flows that have a CDI source.",
		//	        "properties": {
		//	          "CompressionFactor": {
		//	            "description": "A value that is used to calculate compression for an output. The bitrate of the output is calculated as follows: Output bitrate = (1 / compressionFactor) * (source bitrate) This property only applies to outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol. Valid values are in the range of 3.0 to 10.0, inclusive.",
		//	            "type": "number"
		//	          },
		//	          "EncoderProfile": {
		//	            "description": "A setting on the encoder that drives compression settings. This property only applies to video media streams associated with outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol.",
		//	            "enum": [
		//	              "main",
		//	              "high"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "CompressionFactor"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "MediaStreamName": {
		//	        "description": "A name that helps you distinguish one media stream from another.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "EncodingName",
		//	      "MediaStreamName"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"media_stream_output_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DestinationConfigurations
					"destination_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: DestinationIp
								"destination_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The IP address where contents of the media stream will be sent.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: DestinationPort
								"destination_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "The port to use when the content of the media stream is distributed to the output.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Interface
								"interface": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Name
										"name": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The name of the VPC interface that you want to use for the media stream associated with the output.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "The VPC interface that is used for the media stream associated with the output.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "The media streams that you want to associate with the output.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: EncodingName
					"encoding_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The format that will be used to encode the data. For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video streams on sources or outputs that use the CDI protocol, set the encoding name to raw. For video streams on sources or outputs that use the ST 2110 JPEG XS protocol, set the encoding name to jxsv.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: EncodingParameters
					"encoding_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: CompressionFactor
							"compression_factor": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "A value that is used to calculate compression for an output. The bitrate of the output is calculated as follows: Output bitrate = (1 / compressionFactor) * (source bitrate) This property only applies to outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol. Valid values are in the range of 3.0 to 10.0, inclusive.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: EncoderProfile
							"encoder_profile": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A setting on the encoder that drives compression settings. This property only applies to video media streams associated with outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "A collection of parameters that determine how MediaConnect will convert the content. These fields only apply to outputs on flows that have a CDI source.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: MediaStreamName
					"media_stream_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A name that helps you distinguish one media stream from another.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The definition for each media stream that is associated with the output.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MinLatency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The minimum latency in milliseconds.",
		//	  "type": "integer"
		//	}
		"min_latency": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The minimum latency in milliseconds.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the output. This value must be unique within the current flow.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the output. This value must be unique within the current flow.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NdiProgramName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A suffix for the names of the NDI sources that the flow creates. If a custom name isn't specified, MediaConnect uses the output name.",
		//	  "type": "string"
		//	}
		"ndi_program_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A suffix for the names of the NDI sources that the flow creates. If a custom name isn't specified, MediaConnect uses the output name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NdiSpeedHqQuality
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A quality setting for the NDI Speed HQ encoder.",
		//	  "type": "integer"
		//	}
		"ndi_speed_hq_quality": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "A quality setting for the NDI Speed HQ encoder.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutputArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the output.",
		//	  "type": "string"
		//	}
		"output_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the output.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutputStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An indication of whether the output should transmit data or not.",
		//	  "enum": [
		//	    "ENABLED",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"output_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An indication of whether the output should transmit data or not.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Port
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The port to use when content is distributed to this output.",
		//	  "type": "integer"
		//	}
		"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The port to use when content is distributed to this output.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Protocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The protocol that is used by the source or output.",
		//	  "enum": [
		//	    "zixi-push",
		//	    "rtp-fec",
		//	    "rtp",
		//	    "zixi-pull",
		//	    "rist",
		//	    "fujitsu-qos",
		//	    "srt-listener",
		//	    "srt-caller",
		//	    "st2110-jpegxs",
		//	    "cdi",
		//	    "ndi-speed-hq"
		//	  ],
		//	  "type": "string"
		//	}
		"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The protocol that is used by the source or output.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RemoteId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The remote ID for the Zixi-pull stream.",
		//	  "type": "string"
		//	}
		"remote_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The remote ID for the Zixi-pull stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SmoothingLatency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.",
		//	  "type": "integer"
		//	}
		"smoothing_latency": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StreamId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
		//	  "type": "string"
		//	}
		"stream_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcInterfaceAttachment
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The name of the VPC interface attachment to use for this output.",
		//	  "properties": {
		//	    "VpcInterfaceName": {
		//	      "description": "The name of the VPC interface to use for this output.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"vpc_interface_attachment": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: VpcInterfaceName
				"vpc_interface_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the VPC interface to use for this output.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The name of the VPC interface attachment to use for this output.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MediaConnect::FlowOutput",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::FlowOutput").WithTerraformTypeName("awscc_mediaconnect_flow_output")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"algorithm":                          "Algorithm",
		"cidr_allow_list":                    "CidrAllowList",
		"compression_factor":                 "CompressionFactor",
		"description":                        "Description",
		"destination":                        "Destination",
		"destination_configurations":         "DestinationConfigurations",
		"destination_ip":                     "DestinationIp",
		"destination_port":                   "DestinationPort",
		"encoder_profile":                    "EncoderProfile",
		"encoding_name":                      "EncodingName",
		"encoding_parameters":                "EncodingParameters",
		"encryption":                         "Encryption",
		"flow_arn":                           "FlowArn",
		"interface":                          "Interface",
		"key_type":                           "KeyType",
		"max_latency":                        "MaxLatency",
		"media_stream_name":                  "MediaStreamName",
		"media_stream_output_configurations": "MediaStreamOutputConfigurations",
		"min_latency":                        "MinLatency",
		"name":                               "Name",
		"ndi_program_name":                   "NdiProgramName",
		"ndi_speed_hq_quality":               "NdiSpeedHqQuality",
		"output_arn":                         "OutputArn",
		"output_status":                      "OutputStatus",
		"port":                               "Port",
		"protocol":                           "Protocol",
		"remote_id":                          "RemoteId",
		"role_arn":                           "RoleArn",
		"secret_arn":                         "SecretArn",
		"smoothing_latency":                  "SmoothingLatency",
		"stream_id":                          "StreamId",
		"vpc_interface_attachment":           "VpcInterfaceAttachment",
		"vpc_interface_name":                 "VpcInterfaceName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
