---
subcategory: "OpenSearch Serverless"
layout: "aws"
page_title: "AWS: aws_opensearchserverless_collection_group"
description: |-
  Terraform resource for managing an AWS OpenSearch Serverless Collection Group.
---

# Resource: aws_opensearchserverless_collection_group

Manages an AWS OpenSearch Serverless Collection Group.

Collection Groups allow you to manage capacity for multiple collections together, enabling better resource utilization and cost optimization.

## Example Usage

### Basic Usage

```terraform
resource "aws_opensearchserverless_collection_group" "example" {
  name             = "my-collection-group"
  standby_replicas = "ENABLED"
  description      = "Collection group for production logs"
}
```

### With Tags and Description

```terraform
resource "aws_opensearchserverless_collection_group" "example" {
  name             = "my-collection-group"
  standby_replicas = "ENABLED"
  description      = "Collection group for production logs"
  
  tags = {
    Environment = "production"
    Team        = "data-platform"
  }
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the collection group. Must be between 3-32 characters, start with a lowercase letter, and contain only lowercase letters, numbers, and hyphens.
* `standby_replicas` - (Required) Whether standby replicas are enabled. Valid values: `ENABLED`, `DISABLED`.

The following arguments are optional:

* `description` - (Optional) Description of the collection group. Maximum 1000 characters.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Unique identifier for the collection group.
* `arn` - Amazon Resource Name (ARN) of the collection group.
* `created_date` - Unix epoch time when the collection group was created.
* `number_of_collections` - Number of collections associated with the collection group.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `20m`)
* `update` - (Default `20m`)
* `delete` - (Default `20m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import OpenSearch Serverless Collection Group using the `id`. For example:

```terraform
import {
  to = aws_opensearchserverless_collection_group.example
  id = "cg5fzm844zyc2kj7h2e2"
}
```

Using `terraform import`, import OpenSearch Serverless Collection Group using the `id`. For example:

```console
% terraform import aws_opensearchserverless_collection_group.example cg5fzm844zyc2kj7h2e2
```

## Notes

* Collection groups only accept collections of a single type (SEARCH, TIMESERIES, or VECTORSEARCH). The type is determined by the first collection added to the group.
* The `name` and `standby_replicas` fields cannot be changed after creation (ForceNew).
* Capacity limit configuration is not currently supported due to AWS API limitations. This may be added in a future version if AWS adds API support for reading capacity limits.
