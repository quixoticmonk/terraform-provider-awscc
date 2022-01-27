// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package location

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_location_geofence_collection", geofenceCollectionDataSourceType)
}

// geofenceCollectionDataSourceType returns the Terraform awscc_location_geofence_collection data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Location::GeofenceCollection resource type.
func geofenceCollectionDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1600,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"collection_arn": {
			// Property: CollectionArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1600,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"collection_name": {
			// Property: CollectionName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"create_time": {
			// Property: CreateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"pricing_plan": {
			// Property: PricingPlan
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "RequestBasedUsage"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"pricing_plan_data_source": {
			// Property: PricingPlanDataSource
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"update_time": {
			// Property: UpdateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Location::GeofenceCollection",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Location::GeofenceCollection").WithTerraformTypeName("awscc_location_geofence_collection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                      "Arn",
		"collection_arn":           "CollectionArn",
		"collection_name":          "CollectionName",
		"create_time":              "CreateTime",
		"description":              "Description",
		"kms_key_id":               "KmsKeyId",
		"pricing_plan":             "PricingPlan",
		"pricing_plan_data_source": "PricingPlanDataSource",
		"update_time":              "UpdateTime",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
