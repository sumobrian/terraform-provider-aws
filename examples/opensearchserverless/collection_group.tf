# OpenSearch Serverless Collection Group Example

# Create a Collection Group to manage capacity for multiple collections
resource "aws_opensearchserverless_collection_group" "example" {
  name             = "production-logs-group"
  standby_replicas = "ENABLED"
  description      = "Collection group for production log collections"
  
  tags = {
    Environment = "production"
    Team        = "data-platform"
    CostCenter  = "engineering"
  }
}

# Create encryption security policy (required for all collections)
resource "aws_opensearchserverless_security_policy" "example" {
  name        = "production-encryption-policy"
  type        = "encryption"
  description = "Encryption policy for production collections"
  
  policy = jsonencode({
    Rules = [
      {
        ResourceType = "collection"
        Resource     = ["collection/prod-*"]
      }
    ],
    AWSOwnedKey = true
  })
}

# Create collections in the collection group
resource "aws_opensearchserverless_collection" "logs" {
  name                  = "prod-application-logs"
  type                  = "TIMESERIES"
  collection_group_name = aws_opensearchserverless_collection_group.example.name
  description           = "Application logs collection"
  standby_replicas      = "ENABLED"
  
  tags = {
    Environment = "production"
    Application = "web-app"
  }
  
  depends_on = [aws_opensearchserverless_security_policy.example]
}

resource "aws_opensearchserverless_collection" "metrics" {
  name                  = "prod-metrics"
  type                  = "TIMESERIES"
  collection_group_name = aws_opensearchserverless_collection_group.example.name
  description           = "Application metrics collection"
  standby_replicas      = "ENABLED"
  
  tags = {
    Environment = "production"
    Application = "monitoring"
  }
  
  depends_on = [aws_opensearchserverless_security_policy.example]
}

# Query existing collection group
data "aws_opensearchserverless_collection_group" "example" {
  name = "production-logs-group"
}

# Outputs
output "collection_group_id" {
  value       = aws_opensearchserverless_collection_group.example.id
  description = "Collection Group ID"
}

output "collection_group_arn" {
  value       = aws_opensearchserverless_collection_group.example.arn
  description = "Collection Group ARN"
}

output "collections" {
  value = {
    logs = {
      id       = aws_opensearchserverless_collection.logs.id
      endpoint = aws_opensearchserverless_collection.logs.collection_endpoint
    }
    metrics = {
      id       = aws_opensearchserverless_collection.metrics.id
      endpoint = aws_opensearchserverless_collection.metrics.collection_endpoint
    }
  }
  description = "Collection details"
}
