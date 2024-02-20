package dynamodbglobaltable


type DynamodbGlobalTableReplicasTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#key DynamodbGlobalTable#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#value DynamodbGlobalTable#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

