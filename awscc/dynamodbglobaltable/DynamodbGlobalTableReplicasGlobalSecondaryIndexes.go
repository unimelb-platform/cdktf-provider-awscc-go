package dynamodbglobaltable


type DynamodbGlobalTableReplicasGlobalSecondaryIndexes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#index_name DynamodbGlobalTable#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#contributor_insights_specification DynamodbGlobalTable#contributor_insights_specification}.
	ContributorInsightsSpecification *DynamodbGlobalTableReplicasGlobalSecondaryIndexesContributorInsightsSpecification `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#read_provisioned_throughput_settings DynamodbGlobalTable#read_provisioned_throughput_settings}.
	ReadProvisionedThroughputSettings *DynamodbGlobalTableReplicasGlobalSecondaryIndexesReadProvisionedThroughputSettings `field:"optional" json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
}

