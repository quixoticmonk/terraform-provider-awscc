# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an SES email identity
resource "awscc_ses_email_identity" "example" {
  email_identity = "no-reply@example.com"

  # Configure feedback forwarding
  feedback_attributes = {
    email_forwarding_enabled = true
  }
}