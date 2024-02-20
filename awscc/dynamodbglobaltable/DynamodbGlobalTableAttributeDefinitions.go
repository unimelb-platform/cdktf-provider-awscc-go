package dynamodbglobaltable


type DynamodbGlobalTableAttributeDefinitions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#attribute_name DynamodbGlobalTable#attribute_name}.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#attribute_type DynamodbGlobalTable#attribute_type}.
	AttributeType *string `field:"required" json:"attributeType" yaml:"attributeType"`
}

