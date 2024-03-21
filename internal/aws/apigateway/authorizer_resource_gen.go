// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_apigateway_authorizer", authorizerResource)
}

// authorizerResource returns the Terraform awscc_apigateway_authorizer resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGateway::Authorizer resource.
func authorizerResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AuthType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Optional customer-defined field, used in OpenAPI imports and exports without functional impact.",
		//	  "type": "string"
		//	}
		"auth_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Optional customer-defined field, used in OpenAPI imports and exports without functional impact.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerCredentials
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null.",
		//	  "type": "string"
		//	}
		"authorizer_credentials": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"authorizer_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerResultTtlInSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The TTL in seconds of cached authorizer results. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway will cache authorizer responses. If this field is not set, the default value is 300. The maximum value is 3600, or 1 hour.",
		//	  "type": "integer"
		//	}
		"authorizer_result_ttl_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The TTL in seconds of cached authorizer results. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway will cache authorizer responses. If this field is not set, the default value is 300. The maximum value is 3600, or 1 hour.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the authorizer's Uniform Resource Identifier (URI). For ``TOKEN`` or ``REQUEST`` authorizers, this must be a well-formed Lambda function URI, for example, ``arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations``. In general, the URI has this form ``arn:aws:apigateway:{region}:lambda:path/{service_api}``, where ``{region}`` is the same as the region hosting the Lambda function, ``path`` indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial ``/``. For Lambda functions, this is usually of the form ``/2015-03-31/functions/[FunctionARN]/invocations``.",
		//	  "type": "string"
		//	}
		"authorizer_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the authorizer's Uniform Resource Identifier (URI). For ``TOKEN`` or ``REQUEST`` authorizers, this must be a well-formed Lambda function URI, for example, ``arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations``. In general, the URI has this form ``arn:aws:apigateway:{region}:lambda:path/{service_api}``, where ``{region}`` is the same as the region hosting the Lambda function, ``path`` indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial ``/``. For Lambda functions, this is usually of the form ``/2015-03-31/functions/[FunctionARN]/invocations``.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IdentitySource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identity source for which authorization is requested. For a ``TOKEN`` or ``COGNITO_USER_POOLS`` authorizer, this is required and specifies the request header mapping expression for the custom header holding the authorization token submitted by the client. For example, if the token header name is ``Auth``, the header mapping expression is ``method.request.header.Auth``. For the ``REQUEST`` authorizer, this is required when authorization caching is enabled. The value is a comma-separated string of one or more mapping expressions of the specified request parameters. For example, if an ``Auth`` header, a ``Name`` query string parameter are defined as identity sources, this value is ``method.request.header.Auth, method.request.querystring.Name``. These parameters will be used to derive the authorization caching key and to perform runtime validation of the ``REQUEST`` authorizer by verifying all of the identity-related request parameters are present, not null and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function, otherwise, it returns a 401 Unauthorized response without calling the Lambda function. The valid value is a string of comma-separated mapping expressions of the specified request parameters. When the authorization caching is not enabled, this property is optional.",
		//	  "type": "string"
		//	}
		"identity_source": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identity source for which authorization is requested. For a ``TOKEN`` or ``COGNITO_USER_POOLS`` authorizer, this is required and specifies the request header mapping expression for the custom header holding the authorization token submitted by the client. For example, if the token header name is ``Auth``, the header mapping expression is ``method.request.header.Auth``. For the ``REQUEST`` authorizer, this is required when authorization caching is enabled. The value is a comma-separated string of one or more mapping expressions of the specified request parameters. For example, if an ``Auth`` header, a ``Name`` query string parameter are defined as identity sources, this value is ``method.request.header.Auth, method.request.querystring.Name``. These parameters will be used to derive the authorization caching key and to perform runtime validation of the ``REQUEST`` authorizer by verifying all of the identity-related request parameters are present, not null and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function, otherwise, it returns a 401 Unauthorized response without calling the Lambda function. The valid value is a string of comma-separated mapping expressions of the specified request parameters. When the authorization caching is not enabled, this property is optional.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IdentityValidationExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A validation expression for the incoming identity token. For ``TOKEN`` authorizers, this value is a regular expression. For ``COGNITO_USER_POOLS`` authorizers, API Gateway will match the ``aud`` field of the incoming token from the client against the specified regular expression. It will invoke the authorizer's Lambda function when there is a match. Otherwise, it will return a 401 Unauthorized response without calling the Lambda function. The validation expression does not apply to the ``REQUEST`` authorizer.",
		//	  "type": "string"
		//	}
		"identity_validation_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A validation expression for the incoming identity token. For ``TOKEN`` authorizers, this value is a regular expression. For ``COGNITO_USER_POOLS`` authorizers, API Gateway will match the ``aud`` field of the incoming token from the client against the specified regular expression. It will invoke the authorizer's Lambda function when there is a match. Otherwise, it will return a 401 Unauthorized response without calling the Lambda function. The validation expression does not apply to the ``REQUEST`` authorizer.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the authorizer.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the authorizer.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProviderARNs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of the Amazon Cognito user pool ARNs for the ``COGNITO_USER_POOLS`` authorizer. Each element is of this format: ``arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}``. For a ``TOKEN`` or ``REQUEST`` authorizer, this is not defined.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"provider_ar_ns": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of the Amazon Cognito user pool ARNs for the ``COGNITO_USER_POOLS`` authorizer. Each element is of this format: ``arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}``. For a ``TOKEN`` or ``REQUEST`` authorizer, this is not defined.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The string identifier of the associated RestApi.",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The string identifier of the associated RestApi.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The authorizer type. Valid values are ``TOKEN`` for a Lambda function using a single authorization token submitted in a custom header, ``REQUEST`` for a Lambda function using incoming request parameters, and ``COGNITO_USER_POOLS`` for using an Amazon Cognito user pool.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The authorizer type. Valid values are ``TOKEN`` for a Lambda function using a single authorization token submitted in a custom header, ``REQUEST`` for a Lambda function using incoming request parameters, and ``COGNITO_USER_POOLS`` for using an Amazon Cognito user pool.",
			Required:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The ``AWS::ApiGateway::Authorizer`` resource creates an authorization layer that API Gateway activates for methods that have authorization enabled. API Gateway activates the authorizer when a client calls those methods.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Authorizer").WithTerraformTypeName("awscc_apigateway_authorizer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auth_type":                        "AuthType",
		"authorizer_credentials":           "AuthorizerCredentials",
		"authorizer_id":                    "AuthorizerId",
		"authorizer_result_ttl_in_seconds": "AuthorizerResultTtlInSeconds",
		"authorizer_uri":                   "AuthorizerUri",
		"identity_source":                  "IdentitySource",
		"identity_validation_expression":   "IdentityValidationExpression",
		"name":                             "Name",
		"provider_ar_ns":                   "ProviderARNs",
		"rest_api_id":                      "RestApiId",
		"type":                             "Type",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
