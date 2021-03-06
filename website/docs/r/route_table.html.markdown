---
layout: "alicloud"
page_title: "Alicloud: alicloud_route_table"
sidebar_current: "docs-alicloud-resource-route-table"
description: |-
  Provides a Alicloud Route Table resource.
---

# alicloud\_route_table

Provides a route table resource to add customized route tables.

-> **NOTE:** Terraform will auto build route table instance while it uses `alicloud_route_table` to build a route table resource.

Currently, customized route tables are available in most regions apart from China (Beijing), China (Hangzhou), and China (Shenzhen) regions.
For information about route table and how to use it, see [What is Route Table](https://www.alibabacloud.com/help/doc-detail/87057.htm).

## Example Usage

Basic Usage

```
resource "alicloud_route_table" "foo" {
  vpc_id = "vpc-fakeid"
  name = "test_route_table"
  description = "test_route_table"
}
```

## Argument Reference

The following arguments are supported:

* `vpc_id` - (Required, ForceNew) The vpc_id of the route table, the field can't be changed.
* `name` - (Optional) The name of the route table.
* `description` - (Optional) The description of the route table instance.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the route table instance id.

## Import

The route table can be imported using the id, e.g.

```
$ terraform import alicloud_route_table.foo vtb-abc123456
```


