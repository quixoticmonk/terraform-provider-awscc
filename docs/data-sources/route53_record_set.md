---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53_record_set Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Route53::RecordSet
---

# awscc_route53_record_set (Data Source)

Data Source schema for AWS::Route53::RecordSet



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `alias_target` (Attributes) Alias resource record sets only: Information about the AWS resource, such as a CloudFront distribution or an Amazon S3 bucket, that you want to route traffic to. (see [below for nested schema](#nestedatt--alias_target))
- `cidr_routing_config` (Attributes) The object that is specified in resource record set object when you are linking a resource record set to a CIDR location. (see [below for nested schema](#nestedatt--cidr_routing_config))
- `comment` (String) Optional: Any comments you want to include about a change batch request.
- `failover` (String) To configure failover, you add the Failover element to two resource record sets. For one resource record set, you specify PRIMARY as the value for Failover; for the other resource record set, you specify SECONDARY. In addition, you include the HealthCheckId element and specify the health check that you want Amazon Route 53 to perform for each resource record set.
- `geo_location` (Attributes) A complex type that lets you control how Amazon Route 53 responds to DNS queries based on the geographic origin of the query. (see [below for nested schema](#nestedatt--geo_location))
- `health_check_id` (String) If you want Amazon Route 53 to return this resource record set in response to a DNS query only when the status of a health check is healthy, include the HealthCheckId element and specify the ID of the applicable health check.
- `hosted_zone_id` (String) The ID of the hosted zone that you want to create records in.
- `hosted_zone_name` (String) The name of the hosted zone that you want to create records in. You must include a trailing dot (for example, www.example.com.) as part of the HostedZoneName.
- `multi_value_answer` (Boolean) To route traffic approximately randomly to multiple resources, such as web servers, create one multivalue answer record for each resource and specify true for MultiValueAnswer.
- `name` (String) The name of the record that you want to create, update, or delete.
- `region` (String) The Amazon EC2 Region where you created the resource that this resource record set refers to.
- `resource_records` (List of String) One or more values that correspond with the value that you specified for the Type property.
- `set_identifier` (String) An identifier that differentiates among multiple resource record sets that have the same combination of name and type.
- `ttl` (String) The resource record cache time to live (TTL), in seconds.
- `type` (String) The DNS record type.
- `weight` (Number) Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set. Route 53 calculates the sum of the weights for the resource record sets that have the same combination of DNS name and type. Route 53 then responds to queries based on the ratio of a resource's weight to the total.

<a id="nestedatt--alias_target"></a>
### Nested Schema for `alias_target`

Read-Only:

- `dns_name` (String) The value that you specify depends on where you want to route queries.
- `evaluate_target_health` (Boolean) When EvaluateTargetHealth is true, an alias resource record set inherits the health of the referenced AWS resource, such as an ELB load balancer or another resource record set in the hosted zone.
- `hosted_zone_id` (String) The value used depends on where you want to route traffic.


<a id="nestedatt--cidr_routing_config"></a>
### Nested Schema for `cidr_routing_config`

Read-Only:

- `collection_id` (String) The CIDR collection ID.
- `location_name` (String) The CIDR collection location name.


<a id="nestedatt--geo_location"></a>
### Nested Schema for `geo_location`

Read-Only:

- `continent_code` (String) For geolocation resource record sets, a two-letter abbreviation that identifies a continent.
- `country_code` (String) For geolocation resource record sets, the two-letter code for a country.
- `subdivision_code` (String) For geolocation resource record sets, the two-letter code for a state of the United States.