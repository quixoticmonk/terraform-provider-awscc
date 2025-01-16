# Data sources to get the current AWS account ID and region
data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

# Create the Route53 Resolver Firewall Domain List
resource "awscc_route53resolver_firewall_domain_list" "example" {
  name = "example-domain-list"

  # Example of inline domain list
  domains = [
    "example.com",
    "example.org"
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}