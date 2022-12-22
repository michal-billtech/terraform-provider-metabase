---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "metabase_permissions_graph Resource - terraform-provider-metabase"
subcategory: ""
description: |-
  The graph of permissions between permissions groups and databases.
  Metabase exposes a single resource to define all permissions related to databases. This means a single permissions graph resource should be defined in the entire Terraform configuration. However this is not the same as the collection graph, and the two can be combined to grant permissions.
  The permissions graph cannot be created or deleted. Trying to create it will result in an error. It should be imported instead. Trying to delete the resource will succeed with no impact on Metabase (it is a no-op).
  Permissions for the Administrators group cannot be changed. To avoid issues during the update, all permissions for the Administrators group are ignored by default. This behavior can be changed using the ignored groups attribute.
---

# metabase_permissions_graph (Resource)

The graph of permissions between permissions groups and databases.

Metabase exposes a single resource to define all permissions related to databases. This means a single permissions graph resource should be defined in the entire Terraform configuration. However this is not the same as the collection graph, and the two can be combined to grant permissions.

The permissions graph cannot be created or deleted. Trying to create it will result in an error. It should be imported instead. Trying to delete the resource will succeed with no impact on Metabase (it is a no-op).

Permissions for the Administrators group cannot be changed. To avoid issues during the update, all permissions for the Administrators group are ignored by default. This behavior can be changed using the ignored groups attribute.

## Example Usage

```terraform
resource "metabase_database" "bigquery" {
  name = "🗃️ Big Query"

  bigquery_details = {
    service_account_key      = file("sa-key.json")
    project_id               = "gcp-project"
    dataset_filters_type     = "inclusion"
    dataset_filters_patterns = "included_dataset"
  }
}

resource "metabase_permissions_group" "data_analysts" {
  name = "🧑‍🔬 Data Analysts"
}

resource "metabase_permissions_group" "business_stakeholders" {
  name = "👔 Business Stakeholders"
}

resource "metabase_permissions_graph" "graph" {
  advanced_permissions = false

  permissions = [
    {
      group    = metabase_permissions_group.data_analysts.id
      database = metabase_database.bigquery.id
      data = {
        # Native: Yes
        native = "write"
        # Data access: Unrestricted
        schemas = "all"
      }
    },
    {
      group    = metabase_permissions_group.business_stakeholders.id
      database = metabase_database.bigquery.id
      data = {
        # Native: No (by omitting the `native` attribute or setting it to "none")
        # Data access: Unrestricted
        schemas = "all"
      }
    },
    # Permissions for the "All Users" group. Those cannot be removed entirely, but they can be limited.
    # The example below gives the minimum set of permissions for the free version of Metabase:
    {
      group    = 1 # ID for the "All Users" group.
      database = 2
      # Cannot be removed but has no impact when using the free version of Metabase.
      download = {
        native  = "full"
        schemas = "full"
      }
      # Omitting the `data` attribute entirely result in the lowest level of permissions:
      # Data access: No self-service
      # Native: No
    },
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `advanced_permissions` (Boolean) Whether advanced permissions should be set even when not explicitly specified.
- `permissions` (Attributes Set) A list of permissions for a given group and database. A (group, database) pair should appear only once in the list. (see [below for nested schema](#nestedatt--permissions))

### Optional

- `ignored_groups` (Set of Number) The list of group IDs that should be ignored when reading and updating permissions. By default, this contains the Administrators group (`[2]`).

### Read-Only

- `revision` (Number) The revision number for the graph.

<a id="nestedatt--permissions"></a>
### Nested Schema for `permissions`

Required:

- `database` (Number) The ID of the database to which the permission applies.
- `group` (Number) The ID of the group to which the permission applies.

Optional:

- `data` (Attributes) The permission definition for data access. (see [below for nested schema](#nestedatt--permissions--data))
- `data_model` (Attributes) The permission definition for accessing the data model. (see [below for nested schema](#nestedatt--permissions--data_model))
- `details` (String) The permission definition for accessing details.
- `download` (Attributes) The permission definition for downloading data. (see [below for nested schema](#nestedatt--permissions--download))

<a id="nestedatt--permissions--data"></a>
### Nested Schema for `permissions.data`

Optional:

- `native` (String) The permission for native SQL querying
- `schemas` (String) The permission to access data through the Metabase interface


<a id="nestedatt--permissions--data_model"></a>
### Nested Schema for `permissions.data_model`

Optional:

- `native` (String) The permission for native SQL querying
- `schemas` (String) The permission to access data through the Metabase interface


<a id="nestedatt--permissions--download"></a>
### Nested Schema for `permissions.download`

Optional:

- `native` (String) The permission for native SQL querying
- `schemas` (String) The permission to access data through the Metabase interface

## Import

Import is supported using the following syntax:

```shell
# By convention, the revision number of the permissions graph should be used, although it does not really matter as it
# will be read during the import anyway.
terraform import metabase_permissions_graph.graph 1
```