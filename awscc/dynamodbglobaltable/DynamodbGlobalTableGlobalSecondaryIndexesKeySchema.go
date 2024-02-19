package dynamodbglobaltable


type DynamodbGlobalTableGlobalSecondaryIndexesKeySchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#attribute_name DynamodbGlobalTable#attribute_name}.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#key_type DynamodbGlobalTable#key_type}.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
}

