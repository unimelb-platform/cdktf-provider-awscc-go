package dynamodbglobaltable


type DynamodbGlobalTableGlobalSecondaryIndexes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#index_name DynamodbGlobalTable#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#key_schema DynamodbGlobalTable#key_schema}.
	KeySchema interface{} `field:"required" json:"keySchema" yaml:"keySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#projection DynamodbGlobalTable#projection}.
	Projection *DynamodbGlobalTableGlobalSecondaryIndexesProjection `field:"required" json:"projection" yaml:"projection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#write_provisioned_throughput_settings DynamodbGlobalTable#write_provisioned_throughput_settings}.
	WriteProvisionedThroughputSettings *DynamodbGlobalTableGlobalSecondaryIndexesWriteProvisionedThroughputSettings `field:"optional" json:"writeProvisionedThroughputSettings" yaml:"writeProvisionedThroughputSettings"`
}

