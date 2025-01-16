---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_notifications_channel_association Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Notifications::ChannelAssociation
---

# awscc_notifications_channel_association (Data Source)

Data Source schema for AWS::Notifications::ChannelAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) ARN identifier of the channel.
Example: arn:aws:chatbot::123456789012:chat-configuration/slack-channel/security-ops
- `notification_configuration_arn` (String) ARN identifier of the NotificationConfiguration.
Example: arn:aws:notifications::123456789012:configuration/a01jes88qxwkbj05xv9c967pgm1