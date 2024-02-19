package dynamodbtable


type DynamodbTableKinesisStreamSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#stream_arn DynamodbTable#stream_arn}.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#approximate_creation_date_time_precision DynamodbTable#approximate_creation_date_time_precision}.
	ApproximateCreationDateTimePrecision *string `field:"optional" json:"approximateCreationDateTimePrecision" yaml:"approximateCreationDateTimePrecision"`
}

