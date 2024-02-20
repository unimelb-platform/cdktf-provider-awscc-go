package dynamodbglobaltable


type DynamodbGlobalTableWriteProvisionedThroughputSettingsWriteCapacityAutoScalingSettingsTargetTrackingScalingPolicyConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#target_value DynamodbGlobalTable#target_value}.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#disable_scale_in DynamodbGlobalTable#disable_scale_in}.
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#scale_in_cooldown DynamodbGlobalTable#scale_in_cooldown}.
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#scale_out_cooldown DynamodbGlobalTable#scale_out_cooldown}.
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

