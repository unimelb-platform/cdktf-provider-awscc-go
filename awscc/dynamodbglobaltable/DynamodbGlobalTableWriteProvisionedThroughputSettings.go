package dynamodbglobaltable


type DynamodbGlobalTableWriteProvisionedThroughputSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#write_capacity_auto_scaling_settings DynamodbGlobalTable#write_capacity_auto_scaling_settings}.
	WriteCapacityAutoScalingSettings *DynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettings `field:"optional" json:"writeCapacityAutoScalingSettings" yaml:"writeCapacityAutoScalingSettings"`
}

