## 6.31.0 (Unreleased)

FEATURES:

* **New Data Source:** `aws_opensearchserverless_collection_group` ([#46236](https://github.com/hashicorp/terraform-provider-aws/issues/46236))
* **New Ephemeral Resource:** `aws_ecrpublic_authorization_token` ([#45841](https://github.com/hashicorp/terraform-provider-aws/issues/45841))
* **New Resource:** `aws_opensearchserverless_collection_group` ([#46236](https://github.com/hashicorp/terraform-provider-aws/issues/46236))
* **New Resource:** `aws_ssoadmin_customer_managed_policy_attachments_exclusive` ([#46191](https://github.com/hashicorp/terraform-provider-aws/issues/46191))

ENHANCEMENTS:

* resource/aws_odb_cloud_autonomous_vm_cluster: autonomous vm cluster creation using odb network ARN and exadata infrastructure ARN for resource sharing model. ([#45583](https://github.com/hashicorp/terraform-provider-aws/issues/45583))
* resource/aws_opensearch_domain: Add `serverless_vector_acceleration` to `aiml_options` ([#45882](https://github.com/hashicorp/terraform-provider-aws/issues/45882))
* resource/aws_opensearchserverless_collection: Add `collection_group_name` argument to associate collections with collection groups ([#46236](https://github.com/hashicorp/terraform-provider-aws/issues/46236))
* resource/aws_opensearchserverless_collection: Add `encryption_config` configuration block for customer-managed KMS encryption ([#46236](https://github.com/hashicorp/terraform-provider-aws/issues/46236))

BUG FIXES:

* resource/aws_network_interface: Fix `UnauthorizedOperation` error when detaching resource that does not have an attachment ([#46211](https://github.com/hashicorp/terraform-provider-aws/issues/46211))

## 6.30.0 (January 28, 2026)
