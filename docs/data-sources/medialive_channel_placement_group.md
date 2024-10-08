---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_medialive_channel_placement_group Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::MediaLive::ChannelPlacementGroup
---

# awscc_medialive_channel_placement_group (Data Source)

Data Source schema for AWS::MediaLive::ChannelPlacementGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) The ARN of the channel placement group.
- `channel_placement_group_id` (String) Unique internal identifier.
- `channels` (List of String) List of channel IDs added to the channel placement group.
- `cluster_id` (String) The ID of the cluster the node is on.
- `name` (String) The name of the channel placement group.
- `nodes` (List of String) List of nodes added to the channel placement group
- `state` (String) The current state of the ChannelPlacementGroupState
- `tags` (Attributes List) A collection of key-value pairs. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
