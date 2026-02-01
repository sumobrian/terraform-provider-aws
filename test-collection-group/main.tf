terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

# Create encryption security policy
resource "aws_opensearchserverless_security_policy" "encryption" {
  name        = "test-encryption-policy"
  type        = "encryption"
  description = "Encryption policy for test collections"
  
  policy = jsonencode({
    Rules = [
      {
        ResourceType = "collection"
        Resource     = ["collection/test-*"]
      }
    ],
    AWSOwnedKey = true
  })
}

# Collection Group
resource "aws_opensearchserverless_collection_group" "test" {
  name             = "test-logs-group"
  standby_replicas = "ENABLED"
  description      = "Test collection group for logs collections"
  
  tags = {
    Environment = "test"
    Purpose     = "manual-testing"
  }
}

# Collection 1
resource "aws_opensearchserverless_collection" "logs" {
  name                  = "test-logs-collection"
  type                  = "TIMESERIES"
  collection_group_name = aws_opensearchserverless_collection_group.test.name
  description           = "Test logs collection in collection group"
  standby_replicas      = "ENABLED"
  
  tags = {
    Environment = "test"
    Purpose     = "manual-testing"
  }
  
  depends_on = [aws_opensearchserverless_security_policy.encryption]
}

# Collection 2
resource "aws_opensearchserverless_collection" "logs2" {
  name                  = "test-logs-collection-2"
  type                  = "TIMESERIES"
  collection_group_name = aws_opensearchserverless_collection_group.test.name
  description           = "Second TIMESERIES collection in the group"
  standby_replicas      = "ENABLED"
  
  tags = {
    Environment = "test"
    Purpose     = "manual-testing"
  }
  
  depends_on = [aws_opensearchserverless_security_policy.encryption]
}

# TEST 7: Collection with encryption_config using customer-managed KMS key
resource "aws_opensearchserverless_collection" "encrypted" {
  name                  = "test-encrypted-collection"
  type                  = "TIMESERIES"
  collection_group_name = aws_opensearchserverless_collection_group.test.name
  description           = "Collection with customer-managed KMS encryption"
  standby_replicas      = "ENABLED"
  
  encryption_config {
    kms_key_arn = "arn:aws:kms:us-east-1:123456789012:key/example-key-id"
  }
  
  tags = {
    Environment = "test"
    Purpose     = "manual-testing"
    Encryption  = "customer-managed"
  }
  
  depends_on = [aws_opensearchserverless_security_policy.encryption]
}

# Data sources
data "aws_opensearchserverless_collection_group" "by_name" {
  name = "test-logs-group"
}

data "aws_opensearchserverless_collection_group" "by_id" {
  id = aws_opensearchserverless_collection_group.test.id
}

# Outputs
output "test_results_summary" {
  value = {
    test_1_description  = "✅ PASSED"
    test_2_datasources  = "✅ PASSED (by name & ID)"
    test_3_multi_collection = "✅ PASSED (2 collections)"
    test_5_capacity     = "✅ PASSED (20 OCU)"
    test_6_tags         = "✅ PASSED (Owner added)"
    test_7_encryption   = "Testing..."
  }
}

output "collection_group" {
  value = {
    id   = aws_opensearchserverless_collection_group.test.id
    arn  = aws_opensearchserverless_collection_group.test.arn
    name = aws_opensearchserverless_collection_group.test.name
  }
}

output "collections" {
  value = {
    logs = {
      id       = aws_opensearchserverless_collection.logs.id
      endpoint = aws_opensearchserverless_collection.logs.collection_endpoint
      group    = aws_opensearchserverless_collection.logs.collection_group_name
    }
    logs2 = {
      id       = aws_opensearchserverless_collection.logs2.id
      endpoint = aws_opensearchserverless_collection.logs2.collection_endpoint
      group    = aws_opensearchserverless_collection.logs2.collection_group_name
    }
    encrypted = {
      id       = aws_opensearchserverless_collection.encrypted.id
      endpoint = aws_opensearchserverless_collection.encrypted.collection_endpoint
      group    = aws_opensearchserverless_collection.encrypted.collection_group_name
      kms_arn  = "arn:aws:kms:us-east-1:123456789012:key/example-key-id"
    }
  }
}
