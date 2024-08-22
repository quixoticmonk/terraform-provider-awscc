// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package msk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_msk_replicator", replicatorDataSource)
}

// replicatorDataSource returns the Terraform awscc_msk_replicator data source.
// This Terraform data source corresponds to the CloudFormation AWS::MSK::Replicator resource.
func replicatorDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CurrentVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The current version of the MSK replicator.",
		//	  "type": "string"
		//	}
		"current_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The current version of the MSK replicator.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A summary description of the replicator.",
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A summary description of the replicator.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KafkaClusters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies a list of Kafka clusters which are targets of the replicator.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Details of a Kafka cluster for replication.",
		//	    "properties": {
		//	      "AmazonMskCluster": {
		//	        "additionalProperties": false,
		//	        "description": "Details of an Amazon MSK cluster. Exactly one of AmazonMskCluster is required.",
		//	        "properties": {
		//	          "MskClusterArn": {
		//	            "description": "The ARN of an Amazon MSK cluster.",
		//	            "pattern": "arn:(aws|aws-us-gov|aws-cn):kafka:.*",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "MskClusterArn"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "VpcConfig": {
		//	        "additionalProperties": false,
		//	        "description": "Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.",
		//	        "properties": {
		//	          "SecurityGroupIds": {
		//	            "description": "The AWS security groups to associate with the elastic network interfaces in order to specify what the replicator has access to. If a security group is not specified, the default security group associated with the VPC is used.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "maxItems": 16,
		//	            "minItems": 1,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          },
		//	          "SubnetIds": {
		//	            "description": "The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "maxItems": 3,
		//	            "minItems": 2,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          }
		//	        },
		//	        "required": [
		//	          "SubnetIds"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "AmazonMskCluster",
		//	      "VpcConfig"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 2,
		//	  "minItems": 2,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"kafka_clusters": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AmazonMskCluster
					"amazon_msk_cluster": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: MskClusterArn
							"msk_cluster_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The ARN of an Amazon MSK cluster.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Details of an Amazon MSK cluster. Exactly one of AmazonMskCluster is required.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: VpcConfig
					"vpc_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: SecurityGroupIds
							"security_group_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "The AWS security groups to associate with the elastic network interfaces in order to specify what the replicator has access to. If a security group is not specified, the default security group associated with the VPC is used.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: SubnetIds
							"subnet_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Specifies a list of Kafka clusters which are targets of the replicator.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReplicationInfoList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies configuration for replication between a source and target Kafka cluster.",
		//	    "properties": {
		//	      "ConsumerGroupReplication": {
		//	        "additionalProperties": false,
		//	        "description": "Configuration relating to consumer group replication.",
		//	        "properties": {
		//	          "ConsumerGroupsToExclude": {
		//	            "description": "List of regular expression patterns indicating the consumer groups that should not be replicated.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "maxLength": 256,
		//	              "type": "string"
		//	            },
		//	            "maxItems": 100,
		//	            "minItems": 1,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          },
		//	          "ConsumerGroupsToReplicate": {
		//	            "description": "List of regular expression patterns indicating the consumer groups to copy.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "maxLength": 256,
		//	              "type": "string"
		//	            },
		//	            "maxItems": 100,
		//	            "minItems": 0,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          },
		//	          "DetectAndCopyNewConsumerGroups": {
		//	            "description": "Whether to periodically check for new consumer groups.",
		//	            "type": "boolean"
		//	          },
		//	          "SynchroniseConsumerGroupOffsets": {
		//	            "description": "Whether to periodically write the translated offsets to __consumer_offsets topic in target cluster.",
		//	            "type": "boolean"
		//	          }
		//	        },
		//	        "required": [
		//	          "ConsumerGroupsToReplicate"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "SourceKafkaClusterArn": {
		//	        "description": "Amazon Resource Name of the source Kafka cluster.",
		//	        "pattern": "arn:(aws|aws-us-gov|aws-cn):kafka:.*",
		//	        "type": "string"
		//	      },
		//	      "TargetCompressionType": {
		//	        "description": "The type of compression to use writing records to target Kafka cluster.",
		//	        "enum": [
		//	          "NONE",
		//	          "GZIP",
		//	          "SNAPPY",
		//	          "LZ4",
		//	          "ZSTD"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "TargetKafkaClusterArn": {
		//	        "description": "Amazon Resource Name of the target Kafka cluster.",
		//	        "pattern": "arn:(aws|aws-us-gov|aws-cn):kafka:.*",
		//	        "type": "string"
		//	      },
		//	      "TopicReplication": {
		//	        "additionalProperties": false,
		//	        "description": "Configuration relating to topic replication.",
		//	        "properties": {
		//	          "CopyAccessControlListsForTopics": {
		//	            "description": "Whether to periodically configure remote topic ACLs to match their corresponding upstream topics.",
		//	            "type": "boolean"
		//	          },
		//	          "CopyTopicConfigurations": {
		//	            "description": "Whether to periodically configure remote topics to match their corresponding upstream topics.",
		//	            "type": "boolean"
		//	          },
		//	          "DetectAndCopyNewTopics": {
		//	            "description": "Whether to periodically check for new topics and partitions.",
		//	            "type": "boolean"
		//	          },
		//	          "StartingPosition": {
		//	            "additionalProperties": false,
		//	            "description": "Configuration for specifying the position in the topics to start replicating from.",
		//	            "properties": {
		//	              "Type": {
		//	                "description": "The type of replication starting position.",
		//	                "enum": [
		//	                  "LATEST",
		//	                  "EARLIEST"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "TopicNameConfiguration": {
		//	            "additionalProperties": false,
		//	            "description": "Configuration for specifying replicated topic names should be the same as their corresponding upstream topics or prefixed with source cluster alias.",
		//	            "properties": {
		//	              "Type": {
		//	                "description": "The type of replicated topic name.",
		//	                "enum": [
		//	                  "PREFIXED_WITH_SOURCE_CLUSTER_ALIAS",
		//	                  "IDENTICAL"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "TopicsToExclude": {
		//	            "description": "List of regular expression patterns indicating the topics that should not be replicated.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "maxLength": 249,
		//	              "type": "string"
		//	            },
		//	            "maxItems": 100,
		//	            "minItems": 1,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          },
		//	          "TopicsToReplicate": {
		//	            "description": "List of regular expression patterns indicating the topics to copy.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "maxLength": 249,
		//	              "type": "string"
		//	            },
		//	            "maxItems": 100,
		//	            "minItems": 1,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          }
		//	        },
		//	        "required": [
		//	          "TopicsToReplicate"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "SourceKafkaClusterArn",
		//	      "TargetKafkaClusterArn",
		//	      "TopicReplication",
		//	      "ConsumerGroupReplication",
		//	      "TargetCompressionType"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"replication_info_list": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ConsumerGroupReplication
					"consumer_group_replication": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ConsumerGroupsToExclude
							"consumer_groups_to_exclude": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "List of regular expression patterns indicating the consumer groups that should not be replicated.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: ConsumerGroupsToReplicate
							"consumer_groups_to_replicate": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "List of regular expression patterns indicating the consumer groups to copy.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: DetectAndCopyNewConsumerGroups
							"detect_and_copy_new_consumer_groups": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Whether to periodically check for new consumer groups.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: SynchroniseConsumerGroupOffsets
							"synchronise_consumer_group_offsets": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Whether to periodically write the translated offsets to __consumer_offsets topic in target cluster.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Configuration relating to consumer group replication.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: SourceKafkaClusterArn
					"source_kafka_cluster_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Amazon Resource Name of the source Kafka cluster.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: TargetCompressionType
					"target_compression_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of compression to use writing records to target Kafka cluster.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: TargetKafkaClusterArn
					"target_kafka_cluster_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Amazon Resource Name of the target Kafka cluster.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: TopicReplication
					"topic_replication": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: CopyAccessControlListsForTopics
							"copy_access_control_lists_for_topics": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Whether to periodically configure remote topic ACLs to match their corresponding upstream topics.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: CopyTopicConfigurations
							"copy_topic_configurations": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Whether to periodically configure remote topics to match their corresponding upstream topics.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: DetectAndCopyNewTopics
							"detect_and_copy_new_topics": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Whether to periodically check for new topics and partitions.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: StartingPosition
							"starting_position": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Type
									"type": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The type of replication starting position.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Configuration for specifying the position in the topics to start replicating from.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: TopicNameConfiguration
							"topic_name_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Type
									"type": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The type of replicated topic name.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Configuration for specifying replicated topic names should be the same as their corresponding upstream topics or prefixed with source cluster alias.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: TopicsToExclude
							"topics_to_exclude": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "List of regular expression patterns indicating the topics that should not be replicated.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: TopicsToReplicate
							"topics_to_replicate": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "List of regular expression patterns indicating the topics to copy.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Configuration relating to topic replication.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReplicatorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name for the created replicator.",
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn):kafka:.*",
		//	  "type": "string"
		//	}
		"replicator_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name for the created replicator.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReplicatorName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the replicator.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9A-Za-z][0-9A-Za-z-]{0,}$",
		//	  "type": "string"
		//	}
		"replicator_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the replicator.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceExecutionRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.",
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn):iam:.*",
		//	  "type": "string"
		//	}
		"service_execution_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of tags associated with a resource",
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
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
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
			Description: "A collection of tags associated with a resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MSK::Replicator",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MSK::Replicator").WithTerraformTypeName("awscc_msk_replicator")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"amazon_msk_cluster":                   "AmazonMskCluster",
		"consumer_group_replication":           "ConsumerGroupReplication",
		"consumer_groups_to_exclude":           "ConsumerGroupsToExclude",
		"consumer_groups_to_replicate":         "ConsumerGroupsToReplicate",
		"copy_access_control_lists_for_topics": "CopyAccessControlListsForTopics",
		"copy_topic_configurations":            "CopyTopicConfigurations",
		"current_version":                      "CurrentVersion",
		"description":                          "Description",
		"detect_and_copy_new_consumer_groups":  "DetectAndCopyNewConsumerGroups",
		"detect_and_copy_new_topics":           "DetectAndCopyNewTopics",
		"kafka_clusters":                       "KafkaClusters",
		"key":                                  "Key",
		"msk_cluster_arn":                      "MskClusterArn",
		"replication_info_list":                "ReplicationInfoList",
		"replicator_arn":                       "ReplicatorArn",
		"replicator_name":                      "ReplicatorName",
		"security_group_ids":                   "SecurityGroupIds",
		"service_execution_role_arn":           "ServiceExecutionRoleArn",
		"source_kafka_cluster_arn":             "SourceKafkaClusterArn",
		"starting_position":                    "StartingPosition",
		"subnet_ids":                           "SubnetIds",
		"synchronise_consumer_group_offsets":   "SynchroniseConsumerGroupOffsets",
		"tags":                                 "Tags",
		"target_compression_type":              "TargetCompressionType",
		"target_kafka_cluster_arn":             "TargetKafkaClusterArn",
		"topic_name_configuration":             "TopicNameConfiguration",
		"topic_replication":                    "TopicReplication",
		"topics_to_exclude":                    "TopicsToExclude",
		"topics_to_replicate":                  "TopicsToReplicate",
		"type":                                 "Type",
		"value":                                "Value",
		"vpc_config":                           "VpcConfig",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
