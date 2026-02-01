---
subcategory: "OpenSearch Serverless"
layout: "aws"
page_title: "AWS: aws_opensearchserverless_collection_group"
description: |-
  Terraform data source for retrieving an AWS OpenSearch Serverless Collection Group.
---

# Data Source: aws_opensearchserverless_collection_group

Retrieves information about an AWS OpenSearch Serverless Collection Group.

## Example Usage

### By ID

```terraform
data "aws_opensearchserverless_collection_group" "example" {
  id = "cg5fzm844zyc2kj7h2e2"
}
```

### By Name

```terraform
data "aws_opensearchserverless_collection_group" "example" {
  name = "my-collection-group"
}
```

## Argument Reference

~> **NOTE:** One of `id` or `name` must be specified.

The following arguments are optional:

* `id` - (Optional) ID of the collection group.
* `name` - (Optional) Name of the collection group.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the collection group.
* `created_date` - Date the collection group was created (RFC3339 format).
* `description` - Description of the collection group.
* `number_of_collections` - Number of collections associated with the collection group.
* `standby_replicas` - Whether standby replicas are enabled (`ENABLED` or `DISABLED`).
* `tags` - Map of tags assigned to the collection group.
