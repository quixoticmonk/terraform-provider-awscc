// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package pcs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_pcs_compute_node_group", computeNodeGroupDataSource)
}

// computeNodeGroupDataSource returns the Terraform awscc_pcs_compute_node_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::PCS::ComputeNodeGroup resource.
func computeNodeGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AmiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon Machine Image (AMI) that AWS PCS uses to launch instances. If not provided, AWS PCS uses the AMI ID specified in the custom launch template.",
		//	  "pattern": "^ami-[a-z0-9]+$",
		//	  "type": "string"
		//	}
		"ami_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon Machine Image (AMI) that AWS PCS uses to launch instances. If not provided, AWS PCS uses the AMI ID specified in the custom launch template.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique Amazon Resource Name (ARN) of the compute node group.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique Amazon Resource Name (ARN) of the compute node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClusterId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the cluster of the compute node group.",
		//	  "type": "string"
		//	}
		"cluster_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the cluster of the compute node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CustomLaunchTemplate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An Amazon EC2 launch template AWS PCS uses to launch compute nodes.",
		//	  "properties": {
		//	    "TemplateId": {
		//	      "description": "The ID of the EC2 launch template to use to provision instances.",
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "description": "The version of the EC2 launch template to use to provision instances.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Version"
		//	  ],
		//	  "type": "object"
		//	}
		"custom_launch_template": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: TemplateId
				"template_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID of the EC2 launch template to use to provision instances.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The version of the EC2 launch template to use to provision instances.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An Amazon EC2 launch template AWS PCS uses to launch compute nodes.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ErrorInfo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of errors that occurred during compute node group provisioning.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An error that occurred during resource provisioning.",
		//	    "properties": {
		//	      "Code": {
		//	        "description": "The short-form error code.",
		//	        "type": "string"
		//	      },
		//	      "Message": {
		//	        "description": "The detailed error information.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"error_info": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Code
					"code": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The short-form error code.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Message
					"message": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The detailed error information.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of errors that occurred during compute node group provisioning.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IamInstanceProfileArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IAM instance profile used to pass an IAM role when launching EC2 instances. The role contained in your instance profile must have pcs:RegisterComputeNodeGroupInstance permissions attached to provision instances correctly.",
		//	  "pattern": "^arn:aws([a-zA-Z-]{0,10})?:iam::[0-9]{12}:instance-profile/.{1,128}$",
		//	  "type": "string"
		//	}
		"iam_instance_profile_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IAM instance profile used to pass an IAM role when launching EC2 instances. The role contained in your instance profile must have pcs:RegisterComputeNodeGroupInstance permissions attached to provision instances correctly.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The generated unique ID of the compute node group.",
		//	  "type": "string"
		//	}
		"compute_node_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The generated unique ID of the compute node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceConfigs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of EC2 instance configurations that AWS PCS can provision in the compute node group.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An EC2 instance configuration AWS PCS uses to launch compute nodes.",
		//	    "properties": {
		//	      "InstanceType": {
		//	        "description": "The EC2 instance type that AWS PCS can provision in the compute node group.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"instance_configs": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: InstanceType
					"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The EC2 instance type that AWS PCS can provision in the compute node group.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of EC2 instance configurations that AWS PCS can provision in the compute node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name that identifies the compute node group.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name that identifies the compute node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PurchaseOption
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies how EC2 instances are purchased on your behalf. AWS PCS supports On-Demand and Spot instances. For more information, see Instance purchasing options in the Amazon Elastic Compute Cloud User Guide. If you don't provide this option, it defaults to On-Demand.",
		//	  "enum": [
		//	    "ONDEMAND",
		//	    "SPOT"
		//	  ],
		//	  "type": "string"
		//	}
		"purchase_option": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies how EC2 instances are purchased on your behalf. AWS PCS supports On-Demand and Spot instances. For more information, see Instance purchasing options in the Amazon Elastic Compute Cloud User Guide. If you don't provide this option, it defaults to On-Demand.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScalingConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies the boundaries of the compute node group auto scaling.",
		//	  "properties": {
		//	    "MaxInstanceCount": {
		//	      "description": "The upper bound of the number of instances allowed in the compute fleet.",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "MinInstanceCount": {
		//	      "description": "The lower bound of the number of instances allowed in the compute fleet.",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "MaxInstanceCount",
		//	    "MinInstanceCount"
		//	  ],
		//	  "type": "object"
		//	}
		"scaling_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaxInstanceCount
				"max_instance_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The upper bound of the number of instances allowed in the compute fleet.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MinInstanceCount
				"min_instance_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The lower bound of the number of instances allowed in the compute fleet.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies the boundaries of the compute node group auto scaling.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SlurmConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Additional options related to the Slurm scheduler.",
		//	  "properties": {
		//	    "SlurmCustomSettings": {
		//	      "description": "Additional Slurm-specific configuration that directly maps to Slurm settings.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Additional settings that directly map to Slurm settings.",
		//	        "properties": {
		//	          "ParameterName": {
		//	            "description": "AWS PCS supports configuration of the following Slurm parameters for compute node groups: Weight and RealMemory.",
		//	            "type": "string"
		//	          },
		//	          "ParameterValue": {
		//	            "description": "The value for the configured Slurm setting.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "ParameterName",
		//	          "ParameterValue"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"slurm_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SlurmCustomSettings
				"slurm_custom_settings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ParameterName
							"parameter_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "AWS PCS supports configuration of the following Slurm parameters for compute node groups: Weight and RealMemory.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: ParameterValue
							"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The value for the configured Slurm setting.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Additional Slurm-specific configuration that directly maps to Slurm settings.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Additional options related to the Slurm scheduler.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SpotOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Additional configuration when you specify SPOT as the purchase option.",
		//	  "properties": {
		//	    "AllocationStrategy": {
		//	      "description": "The Amazon EC2 allocation strategy AWS PCS uses to provision EC2 instances. AWS PCS supports lowest price, capacity optimized, and price capacity optimized. If you don't provide this option, it defaults to price capacity optimized.",
		//	      "enum": [
		//	        "lowest-price",
		//	        "capacity-optimized",
		//	        "price-capacity-optimized"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"spot_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllocationStrategy
				"allocation_strategy": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon EC2 allocation strategy AWS PCS uses to provision EC2 instances. AWS PCS supports lowest price, capacity optimized, and price capacity optimized. If you don't provide this option, it defaults to price capacity optimized.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Additional configuration when you specify SPOT as the purchase option.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The provisioning status of the compute node group. The provisioning status doesn't indicate the overall health of the compute node group.",
		//	  "enum": [
		//	    "CREATING",
		//	    "ACTIVE",
		//	    "UPDATING",
		//	    "DELETING",
		//	    "CREATE_FAILED",
		//	    "DELETE_FAILED",
		//	    "UPDATE_FAILED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The provisioning status of the compute node group. The provisioning status doesn't indicate the overall health of the compute node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of subnet IDs where instances are provisioned by the compute node group. The subnets must be in the same VPC as the cluster.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "A VPC subnet ID.",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The list of subnet IDs where instances are provisioned by the compute node group. The subnets must be in the same VPC as the cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  }
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::PCS::ComputeNodeGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::PCS::ComputeNodeGroup").WithTerraformTypeName("awscc_pcs_compute_node_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocation_strategy":      "AllocationStrategy",
		"ami_id":                   "AmiId",
		"arn":                      "Arn",
		"cluster_id":               "ClusterId",
		"code":                     "Code",
		"compute_node_group_id":    "Id",
		"custom_launch_template":   "CustomLaunchTemplate",
		"error_info":               "ErrorInfo",
		"iam_instance_profile_arn": "IamInstanceProfileArn",
		"instance_configs":         "InstanceConfigs",
		"instance_type":            "InstanceType",
		"max_instance_count":       "MaxInstanceCount",
		"message":                  "Message",
		"min_instance_count":       "MinInstanceCount",
		"name":                     "Name",
		"parameter_name":           "ParameterName",
		"parameter_value":          "ParameterValue",
		"purchase_option":          "PurchaseOption",
		"scaling_configuration":    "ScalingConfiguration",
		"slurm_configuration":      "SlurmConfiguration",
		"slurm_custom_settings":    "SlurmCustomSettings",
		"spot_options":             "SpotOptions",
		"status":                   "Status",
		"subnet_ids":               "SubnetIds",
		"tags":                     "Tags",
		"template_id":              "TemplateId",
		"version":                  "Version",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
