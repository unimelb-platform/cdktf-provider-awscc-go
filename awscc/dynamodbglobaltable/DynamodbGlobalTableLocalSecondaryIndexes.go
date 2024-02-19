package dynamodbglobaltable


type DynamodbGlobalTableLocalSecondaryIndexes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#index_name DynamodbGlobalTable#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#key_schema DynamodbGlobalTable#key_schema}.
	KeySchema interface{} `field:"required" json:"keySchema" yaml:"keySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#projection DynamodbGlobalTable#projection}.
	Projection *DynamodbGlobalTableLocalSecondaryIndexesProjection `field:"required" json:"projection" yaml:"projection"`
}

