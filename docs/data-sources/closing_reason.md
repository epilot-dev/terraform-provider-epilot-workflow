---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "epilot-workflow_closing_reason Data Source - terraform-provider-epilot-workflow"
subcategory: ""
description: |-
  ClosingReason DataSource
---

# epilot-workflow_closing_reason (Data Source)

ClosingReason DataSource

## Example Usage

```terraform
data "epilot-workflow_closing_reason" "my_closingreason" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `creation_time` (String)
- `id` (String) The ID of this resource.
- `last_update_time` (String)
- `status` (String)
- `title` (String)
