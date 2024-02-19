package dynamodbtable


type DynamodbTableGlobalSecondaryIndexesKeySchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#attribute_name DynamodbTable#attribute_name}.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#key_type DynamodbTable#key_type}.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
}

