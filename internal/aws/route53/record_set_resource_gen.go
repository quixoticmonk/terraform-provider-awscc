// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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
	registry.AddResourceFactory("awscc_route53_record_set", recordSetResource)
}

// recordSetResource returns the Terraform awscc_route53_record_set resource.
// This Terraform resource corresponds to the CloudFormation AWS::Route53::RecordSet resource.
func recordSetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AliasTarget
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Alias resource record sets only: Information about the AWS resource, such as a CloudFront distribution or an Amazon S3 bucket, that you want to route traffic to.",
		//	  "properties": {
		//	    "DNSName": {
		//	      "description": "The value that you specify depends on where you want to route queries.",
		//	      "maxLength": 1024,
		//	      "type": "string"
		//	    },
		//	    "EvaluateTargetHealth": {
		//	      "default": false,
		//	      "description": "When EvaluateTargetHealth is true, an alias resource record set inherits the health of the referenced AWS resource, such as an ELB load balancer or another resource record set in the hosted zone.",
		//	      "type": "boolean"
		//	    },
		//	    "HostedZoneId": {
		//	      "description": "The value used depends on where you want to route traffic.",
		//	      "maxLength": 44,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "HostedZoneId",
		//	    "DNSName"
		//	  ],
		//	  "type": "object"
		//	}
		"alias_target": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DNSName
				"dns_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The value that you specify depends on where you want to route queries.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(1024),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: EvaluateTargetHealth
				"evaluate_target_health": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "When EvaluateTargetHealth is true, an alias resource record set inherits the health of the referenced AWS resource, such as an ELB load balancer or another resource record set in the hosted zone.",
					Optional:    true,
					Computed:    true,
					Default:     booldefault.StaticBool(false),
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: HostedZoneId
				"hosted_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The value used depends on where you want to route traffic.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(44),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Alias resource record sets only: Information about the AWS resource, such as a CloudFront distribution or an Amazon S3 bucket, that you want to route traffic to.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CidrRoutingConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The object that is specified in resource record set object when you are linking a resource record set to a CIDR location.",
		//	  "properties": {
		//	    "CollectionId": {
		//	      "description": "The CIDR collection ID.",
		//	      "pattern": "^[0-9a-f]{8}-(?:[0-9a-f]{4}-){3}[0-9a-f]{12}$",
		//	      "type": "string"
		//	    },
		//	    "LocationName": {
		//	      "description": "The CIDR collection location name.",
		//	      "maxLength": 16,
		//	      "minLength": 1,
		//	      "pattern": "[0-9A-Za-z_\\-\\*]+",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "CollectionId",
		//	    "LocationName"
		//	  ],
		//	  "type": "object"
		//	}
		"cidr_routing_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CollectionId
				"collection_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The CIDR collection ID.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-f]{8}-(?:[0-9a-f]{4}-){3}[0-9a-f]{12}$"), ""),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: LocationName
				"location_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The CIDR collection location name.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 16),
						stringvalidator.RegexMatches(regexp.MustCompile("[0-9A-Za-z_\\-\\*]+"), ""),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The object that is specified in resource record set object when you are linking a resource record set to a CIDR location.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Comment
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Optional: Any comments you want to include about a change batch request.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Optional: Any comments you want to include about a change batch request.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Comment is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Failover
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "To configure failover, you add the Failover element to two resource record sets. For one resource record set, you specify PRIMARY as the value for Failover; for the other resource record set, you specify SECONDARY. In addition, you include the HealthCheckId element and specify the health check that you want Amazon Route 53 to perform for each resource record set.",
		//	  "enum": [
		//	    "PRIMARY",
		//	    "SECONDARY"
		//	  ],
		//	  "type": "string"
		//	}
		"failover": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "To configure failover, you add the Failover element to two resource record sets. For one resource record set, you specify PRIMARY as the value for Failover; for the other resource record set, you specify SECONDARY. In addition, you include the HealthCheckId element and specify the health check that you want Amazon Route 53 to perform for each resource record set.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"PRIMARY",
					"SECONDARY",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GeoLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A complex type that lets you control how Amazon Route 53 responds to DNS queries based on the geographic origin of the query.",
		//	  "oneOf": [
		//	    {
		//	      "required": [
		//	        "ContinentCode"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "CountryCode"
		//	      ]
		//	    }
		//	  ],
		//	  "properties": {
		//	    "ContinentCode": {
		//	      "description": "For geolocation resource record sets, a two-letter abbreviation that identifies a continent.",
		//	      "maxLength": 2,
		//	      "minLength": 2,
		//	      "type": "string"
		//	    },
		//	    "CountryCode": {
		//	      "description": "For geolocation resource record sets, the two-letter code for a country.",
		//	      "maxLength": 2,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "SubdivisionCode": {
		//	      "description": "For geolocation resource record sets, the two-letter code for a state of the United States.",
		//	      "maxLength": 3,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"geo_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ContinentCode
				"continent_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "For geolocation resource record sets, a two-letter abbreviation that identifies a continent.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(2, 2),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CountryCode
				"country_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "For geolocation resource record sets, the two-letter code for a country.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 2),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SubdivisionCode
				"subdivision_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "For geolocation resource record sets, the two-letter code for a state of the United States.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 3),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A complex type that lets you control how Amazon Route 53 responds to DNS queries based on the geographic origin of the query.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If you want Amazon Route 53 to return this resource record set in response to a DNS query only when the status of a health check is healthy, include the HealthCheckId element and specify the ID of the applicable health check.",
		//	  "maxLength": 64,
		//	  "type": "string"
		//	}
		"health_check_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "If you want Amazon Route 53 to return this resource record set in response to a DNS query only when the status of a health check is healthy, include the HealthCheckId element and specify the ID of the applicable health check.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HostedZoneId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the hosted zone that you want to create records in.",
		//	  "maxLength": 44,
		//	  "type": "string"
		//	}
		"hosted_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the hosted zone that you want to create records in.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(44),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HostedZoneName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the hosted zone that you want to create records in. You must include a trailing dot (for example, www.example.com.) as part of the HostedZoneName.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"hosted_zone_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the hosted zone that you want to create records in. You must include a trailing dot (for example, www.example.com.) as part of the HostedZoneName.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// HostedZoneName is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: MultiValueAnswer
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "To route traffic approximately randomly to multiple resources, such as web servers, create one multivalue answer record for each resource and specify true for MultiValueAnswer.",
		//	  "type": "boolean"
		//	}
		"multi_value_answer": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "To route traffic approximately randomly to multiple resources, such as web servers, create one multivalue answer record for each resource and specify true for MultiValueAnswer.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the record that you want to create, update, or delete.",
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the record that you want to create, update, or delete.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1024),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Region
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon EC2 Region where you created the resource that this resource record set refers to.",
		//	  "type": "string"
		//	}
		"region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon EC2 Region where you created the resource that this resource record set refers to.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceRecords
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more values that correspond with the value that you specified for the Type property.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "maxLength": 4000,
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"resource_records": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "One or more values that correspond with the value that you specified for the Type property.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthAtMost(4000),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SetIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An identifier that differentiates among multiple resource record sets that have the same combination of name and type.",
		//	  "maxLength": 128,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"set_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An identifier that differentiates among multiple resource record sets that have the same combination of name and type.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 128),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TTL
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The resource record cache time to live (TTL), in seconds.",
		//	  "type": "string"
		//	}
		"ttl": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The resource record cache time to live (TTL), in seconds.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The DNS record type.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The DNS record type.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Weight
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set. Route 53 calculates the sum of the weights for the resource record sets that have the same combination of DNS name and type. Route 53 then responds to queries based on the ratio of a resource's weight to the total.",
		//	  "type": "integer"
		//	}
		"weight": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set. Route 53 calculates the sum of the weights for the resource record sets that have the same combination of DNS name and type. Route 53 then responds to queries based on the ratio of a resource's weight to the total.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::Route53::RecordSet.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53::RecordSet").WithTerraformTypeName("awscc_route53_record_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias_target":           "AliasTarget",
		"cidr_routing_config":    "CidrRoutingConfig",
		"collection_id":          "CollectionId",
		"comment":                "Comment",
		"continent_code":         "ContinentCode",
		"country_code":           "CountryCode",
		"dns_name":               "DNSName",
		"evaluate_target_health": "EvaluateTargetHealth",
		"failover":               "Failover",
		"geo_location":           "GeoLocation",
		"health_check_id":        "HealthCheckId",
		"hosted_zone_id":         "HostedZoneId",
		"hosted_zone_name":       "HostedZoneName",
		"location_name":          "LocationName",
		"multi_value_answer":     "MultiValueAnswer",
		"name":                   "Name",
		"region":                 "Region",
		"resource_records":       "ResourceRecords",
		"set_identifier":         "SetIdentifier",
		"subdivision_code":       "SubdivisionCode",
		"ttl":                    "TTL",
		"type":                   "Type",
		"weight":                 "Weight",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Comment",
		"/properties/HostedZoneName",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}