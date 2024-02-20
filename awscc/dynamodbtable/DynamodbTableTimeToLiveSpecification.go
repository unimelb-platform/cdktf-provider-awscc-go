package dynamodbtable


type DynamodbTableTimeToLiveSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#enabled DynamodbTable#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#attribute_name DynamodbTable#attribute_name}.
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
}

