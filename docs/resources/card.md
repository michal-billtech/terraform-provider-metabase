---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "metabase_card Resource - terraform-provider-metabase"
subcategory: ""
description: |-
  A Metabase card (question).
  Because the content of a card is complex and can vary a lot between cards, the full schema is not defined in Terraform, and a JSON string should be used instead. You can use templatefile or jsonencode to make the experience smoother.
---

# metabase_card (Resource)

A Metabase card (question).

Because the content of a card is complex and can vary a lot between cards, the full schema is not defined in Terraform, and a JSON string should be used instead. You can use templatefile or jsonencode to make the experience smoother.

## Example Usage

```terraform
data "metabase_table" "table" {
  name = "my_table"
}

resource "metabase_card" "some_great_insights" {
  # To avoid unnecessary diffs, the exact list of attributes below should be specified.
  json = jsonencode({
    name                = "💡 Some great insights"
    description         = "📖 You will learn a lot from this graph!"
    collection_id       = null
    collection_position = null
    cache_ttl           = null
    query_type          = "query"
    dataset_query = {
      database = data.metabase_table.table.db_id
      query = {
        source-table = data.metabase_table.table.id
        aggregation = [
          ["count"]
        ]
        breakout = [
          ["field", data.metabase_table.table.fields["group_by_column"], null]
        ]
      }
      type = "query"
    }
    parameter_mappings     = []
    display                = "pie"
    visualization_settings = {}
    parameters             = []
  })
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `json` (String) The full card definition as a JSON string.

### Read-Only

- `id` (Number) The ID of the card.

## Import

Import is supported using the following syntax:

```shell
# Use the integer ID from the Metabase API.
terraform import metabase_card.card 1
```