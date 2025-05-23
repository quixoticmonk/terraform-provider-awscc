---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_bedrock_data_automation_project Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Bedrock::DataAutomationProject
---

# awscc_bedrock_data_automation_project (Data Source)

Data Source schema for AWS::Bedrock::DataAutomationProject



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `creation_time` (String) Time Stamp
- `custom_output_configuration` (Attributes) Custom output configuration (see [below for nested schema](#nestedatt--custom_output_configuration))
- `kms_encryption_context` (Map of String) KMS encryption context
- `kms_key_id` (String) KMS key identifier
- `last_modified_time` (String) Time Stamp
- `override_configuration` (Attributes) Override configuration (see [below for nested schema](#nestedatt--override_configuration))
- `project_arn` (String) ARN of a DataAutomationProject
- `project_description` (String) Description of the DataAutomationProject
- `project_name` (String) Name of the DataAutomationProject
- `project_stage` (String) Stage of the Project
- `standard_output_configuration` (Attributes) Standard output configuration (see [below for nested schema](#nestedatt--standard_output_configuration))
- `status` (String)
- `tags` (Attributes List) List of Tags (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--custom_output_configuration"></a>
### Nested Schema for `custom_output_configuration`

Read-Only:

- `blueprints` (Attributes List) (see [below for nested schema](#nestedatt--custom_output_configuration--blueprints))

<a id="nestedatt--custom_output_configuration--blueprints"></a>
### Nested Schema for `custom_output_configuration.blueprints`

Read-Only:

- `blueprint_arn` (String) ARN of a Blueprint
- `blueprint_stage` (String) Stage of the Blueprint
- `blueprint_version` (String) Blueprint Version



<a id="nestedatt--override_configuration"></a>
### Nested Schema for `override_configuration`

Read-Only:

- `audio` (Attributes) (see [below for nested schema](#nestedatt--override_configuration--audio))
- `document` (Attributes) (see [below for nested schema](#nestedatt--override_configuration--document))
- `image` (Attributes) (see [below for nested schema](#nestedatt--override_configuration--image))
- `modality_routing` (Attributes) Modality routing configuration (see [below for nested schema](#nestedatt--override_configuration--modality_routing))
- `video` (Attributes) (see [below for nested schema](#nestedatt--override_configuration--video))

<a id="nestedatt--override_configuration--audio"></a>
### Nested Schema for `override_configuration.audio`

Read-Only:

- `modality_processing` (Attributes) (see [below for nested schema](#nestedatt--override_configuration--audio--modality_processing))

<a id="nestedatt--override_configuration--audio--modality_processing"></a>
### Nested Schema for `override_configuration.audio.modality_processing`

Read-Only:

- `state` (String)



<a id="nestedatt--override_configuration--document"></a>
### Nested Schema for `override_configuration.document`

Read-Only:

- `modality_processing` (Attributes) (see [below for nested schema](#nestedatt--override_configuration--document--modality_processing))
- `splitter` (Attributes) (see [below for nested schema](#nestedatt--override_configuration--document--splitter))

<a id="nestedatt--override_configuration--document--modality_processing"></a>
### Nested Schema for `override_configuration.document.modality_processing`

Read-Only:

- `state` (String)


<a id="nestedatt--override_configuration--document--splitter"></a>
### Nested Schema for `override_configuration.document.splitter`

Read-Only:

- `state` (String)



<a id="nestedatt--override_configuration--image"></a>
### Nested Schema for `override_configuration.image`

Read-Only:

- `modality_processing` (Attributes) (see [below for nested schema](#nestedatt--override_configuration--image--modality_processing))

<a id="nestedatt--override_configuration--image--modality_processing"></a>
### Nested Schema for `override_configuration.image.modality_processing`

Read-Only:

- `state` (String)



<a id="nestedatt--override_configuration--modality_routing"></a>
### Nested Schema for `override_configuration.modality_routing`

Read-Only:

- `jpeg` (String)
- `mov` (String)
- `mp_4` (String)
- `png` (String)


<a id="nestedatt--override_configuration--video"></a>
### Nested Schema for `override_configuration.video`

Read-Only:

- `modality_processing` (Attributes) (see [below for nested schema](#nestedatt--override_configuration--video--modality_processing))

<a id="nestedatt--override_configuration--video--modality_processing"></a>
### Nested Schema for `override_configuration.video.modality_processing`

Read-Only:

- `state` (String)




<a id="nestedatt--standard_output_configuration"></a>
### Nested Schema for `standard_output_configuration`

Read-Only:

- `audio` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--audio))
- `document` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--document))
- `image` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--image))
- `video` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--video))

<a id="nestedatt--standard_output_configuration--audio"></a>
### Nested Schema for `standard_output_configuration.audio`

Read-Only:

- `extraction` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--audio--extraction))
- `generative_field` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--audio--generative_field))

<a id="nestedatt--standard_output_configuration--audio--extraction"></a>
### Nested Schema for `standard_output_configuration.audio.extraction`

Read-Only:

- `category` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--audio--extraction--category))

<a id="nestedatt--standard_output_configuration--audio--extraction--category"></a>
### Nested Schema for `standard_output_configuration.audio.extraction.category`

Read-Only:

- `state` (String)
- `types` (List of String)



<a id="nestedatt--standard_output_configuration--audio--generative_field"></a>
### Nested Schema for `standard_output_configuration.audio.generative_field`

Read-Only:

- `state` (String)
- `types` (List of String)



<a id="nestedatt--standard_output_configuration--document"></a>
### Nested Schema for `standard_output_configuration.document`

Read-Only:

- `extraction` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--document--extraction))
- `generative_field` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--document--generative_field))
- `output_format` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--document--output_format))

<a id="nestedatt--standard_output_configuration--document--extraction"></a>
### Nested Schema for `standard_output_configuration.document.extraction`

Read-Only:

- `bounding_box` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--document--extraction--bounding_box))
- `granularity` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--document--extraction--granularity))

<a id="nestedatt--standard_output_configuration--document--extraction--bounding_box"></a>
### Nested Schema for `standard_output_configuration.document.extraction.bounding_box`

Read-Only:

- `state` (String)


<a id="nestedatt--standard_output_configuration--document--extraction--granularity"></a>
### Nested Schema for `standard_output_configuration.document.extraction.granularity`

Read-Only:

- `types` (List of String)



<a id="nestedatt--standard_output_configuration--document--generative_field"></a>
### Nested Schema for `standard_output_configuration.document.generative_field`

Read-Only:

- `state` (String)


<a id="nestedatt--standard_output_configuration--document--output_format"></a>
### Nested Schema for `standard_output_configuration.document.output_format`

Read-Only:

- `additional_file_format` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--document--output_format--additional_file_format))
- `text_format` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--document--output_format--text_format))

<a id="nestedatt--standard_output_configuration--document--output_format--additional_file_format"></a>
### Nested Schema for `standard_output_configuration.document.output_format.additional_file_format`

Read-Only:

- `state` (String)


<a id="nestedatt--standard_output_configuration--document--output_format--text_format"></a>
### Nested Schema for `standard_output_configuration.document.output_format.text_format`

Read-Only:

- `types` (List of String)




<a id="nestedatt--standard_output_configuration--image"></a>
### Nested Schema for `standard_output_configuration.image`

Read-Only:

- `extraction` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--image--extraction))
- `generative_field` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--image--generative_field))

<a id="nestedatt--standard_output_configuration--image--extraction"></a>
### Nested Schema for `standard_output_configuration.image.extraction`

Read-Only:

- `bounding_box` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--image--extraction--bounding_box))
- `category` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--image--extraction--category))

<a id="nestedatt--standard_output_configuration--image--extraction--bounding_box"></a>
### Nested Schema for `standard_output_configuration.image.extraction.bounding_box`

Read-Only:

- `state` (String)


<a id="nestedatt--standard_output_configuration--image--extraction--category"></a>
### Nested Schema for `standard_output_configuration.image.extraction.category`

Read-Only:

- `state` (String)
- `types` (List of String)



<a id="nestedatt--standard_output_configuration--image--generative_field"></a>
### Nested Schema for `standard_output_configuration.image.generative_field`

Read-Only:

- `state` (String)
- `types` (List of String)



<a id="nestedatt--standard_output_configuration--video"></a>
### Nested Schema for `standard_output_configuration.video`

Read-Only:

- `extraction` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--video--extraction))
- `generative_field` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--video--generative_field))

<a id="nestedatt--standard_output_configuration--video--extraction"></a>
### Nested Schema for `standard_output_configuration.video.extraction`

Read-Only:

- `bounding_box` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--video--extraction--bounding_box))
- `category` (Attributes) (see [below for nested schema](#nestedatt--standard_output_configuration--video--extraction--category))

<a id="nestedatt--standard_output_configuration--video--extraction--bounding_box"></a>
### Nested Schema for `standard_output_configuration.video.extraction.bounding_box`

Read-Only:

- `state` (String)


<a id="nestedatt--standard_output_configuration--video--extraction--category"></a>
### Nested Schema for `standard_output_configuration.video.extraction.category`

Read-Only:

- `state` (String)
- `types` (List of String)



<a id="nestedatt--standard_output_configuration--video--generative_field"></a>
### Nested Schema for `standard_output_configuration.video.generative_field`

Read-Only:

- `state` (String)
- `types` (List of String)




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) Key for the tag
- `value` (String) Value for the tag
