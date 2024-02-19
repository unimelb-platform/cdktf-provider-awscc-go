package dynamodbtable


type DynamodbTableAttributeDefinitions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#attribute_name DynamodbTable#attribute_name}.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#attribute_type DynamodbTable#attribute_type}.
	AttributeType *string `field:"required" json:"attributeType" yaml:"attributeType"`
}

