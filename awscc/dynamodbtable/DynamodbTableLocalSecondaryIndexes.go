package dynamodbtable


type DynamodbTableLocalSecondaryIndexes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#index_name DynamodbTable#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#key_schema DynamodbTable#key_schema}.
	KeySchema interface{} `field:"required" json:"keySchema" yaml:"keySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#projection DynamodbTable#projection}.
	Projection *DynamodbTableLocalSecondaryIndexesProjection `field:"required" json:"projection" yaml:"projection"`
}

