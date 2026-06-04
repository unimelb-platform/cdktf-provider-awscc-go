package dynamodbglobaltable


type DynamodbGlobalTableWarmThroughput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dynamodb_global_table#read_units_per_second DynamodbGlobalTable#read_units_per_second}.
	ReadUnitsPerSecond *float64 `field:"optional" json:"readUnitsPerSecond" yaml:"readUnitsPerSecond"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dynamodb_global_table#write_units_per_second DynamodbGlobalTable#write_units_per_second}.
	WriteUnitsPerSecond *float64 `field:"optional" json:"writeUnitsPerSecond" yaml:"writeUnitsPerSecond"`
}

