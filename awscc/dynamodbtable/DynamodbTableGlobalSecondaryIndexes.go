package dynamodbtable


type DynamodbTableGlobalSecondaryIndexes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#index_name DynamodbTable#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#key_schema DynamodbTable#key_schema}.
	KeySchema interface{} `field:"required" json:"keySchema" yaml:"keySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#projection DynamodbTable#projection}.
	Projection *DynamodbTableGlobalSecondaryIndexesProjection `field:"required" json:"projection" yaml:"projection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#contributor_insights_specification DynamodbTable#contributor_insights_specification}.
	ContributorInsightsSpecification *DynamodbTableGlobalSecondaryIndexesContributorInsightsSpecification `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#provisioned_throughput DynamodbTable#provisioned_throughput}.
	ProvisionedThroughput *DynamodbTableGlobalSecondaryIndexesProvisionedThroughput `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
}

