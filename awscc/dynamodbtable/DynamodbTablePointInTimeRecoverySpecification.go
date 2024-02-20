package dynamodbtable


type DynamodbTablePointInTimeRecoverySpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#point_in_time_recovery_enabled DynamodbTable#point_in_time_recovery_enabled}.
	PointInTimeRecoveryEnabled interface{} `field:"optional" json:"pointInTimeRecoveryEnabled" yaml:"pointInTimeRecoveryEnabled"`
}

