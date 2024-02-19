package dynamodbglobaltable


type DynamodbGlobalTableReplicasReadProvisionedThroughputSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#read_capacity_auto_scaling_settings DynamodbGlobalTable#read_capacity_auto_scaling_settings}.
	ReadCapacityAutoScalingSettings *DynamodbGlobalTableReplicasReadProvisionedThroughputSettingsReadCapacityAutoScalingSettings `field:"optional" json:"readCapacityAutoScalingSettings" yaml:"readCapacityAutoScalingSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#read_capacity_units DynamodbGlobalTable#read_capacity_units}.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
}

