package dynamodbglobaltable


type DynamodbGlobalTableTimeToLiveSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#enabled DynamodbGlobalTable#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#attribute_name DynamodbGlobalTable#attribute_name}.
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
}

