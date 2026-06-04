package dynamodbglobaltable


type DynamodbGlobalTableWriteOnDemandThroughputSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dynamodb_global_table#max_write_request_units DynamodbGlobalTable#max_write_request_units}.
	MaxWriteRequestUnits *float64 `field:"optional" json:"maxWriteRequestUnits" yaml:"maxWriteRequestUnits"`
}

