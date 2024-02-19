package dynamodbglobaltable


type DynamodbGlobalTableLocalSecondaryIndexesProjection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#non_key_attributes DynamodbGlobalTable#non_key_attributes}.
	NonKeyAttributes *[]*string `field:"optional" json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#projection_type DynamodbGlobalTable#projection_type}.
	ProjectionType *string `field:"optional" json:"projectionType" yaml:"projectionType"`
}

