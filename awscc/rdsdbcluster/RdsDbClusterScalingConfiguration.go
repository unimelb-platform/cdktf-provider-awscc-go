package rdsdbcluster


type RdsDbClusterScalingConfiguration struct {
	// A value that indicates whether to allow or disallow automatic pause for an Aurora DB cluster in serverless DB engine mode.
	//
	// A DB cluster can be paused only when it's idle (it has no connections).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#auto_pause RdsDbCluster#auto_pause}
	AutoPause interface{} `field:"optional" json:"autoPause" yaml:"autoPause"`
	// The maximum capacity for an Aurora DB cluster in serverless DB engine mode.
	//
	// For Aurora MySQL, valid capacity values are 1, 2, 4, 8, 16, 32, 64, 128, and 256.
	// For Aurora PostgreSQL, valid capacity values are 2, 4, 8, 16, 32, 64, 192, and 384.
	// The maximum capacity must be greater than or equal to the minimum capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#max_capacity RdsDbCluster#max_capacity}
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity for an Aurora DB cluster in serverless DB engine mode.
	//
	// For Aurora MySQL, valid capacity values are 1, 2, 4, 8, 16, 32, 64, 128, and 256.
	// For Aurora PostgreSQL, valid capacity values are 2, 4, 8, 16, 32, 64, 192, and 384.
	// The minimum capacity must be less than or equal to the maximum capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#min_capacity RdsDbCluster#min_capacity}
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// The amount of time, in seconds, that Aurora Serverless v1 tries to find a scaling point to perform seamless scaling before enforcing the timeout action.
	//
	// The default is 300.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#seconds_before_timeout RdsDbCluster#seconds_before_timeout}
	SecondsBeforeTimeout *float64 `field:"optional" json:"secondsBeforeTimeout" yaml:"secondsBeforeTimeout"`
	// The time, in seconds, before an Aurora DB cluster in serverless mode is paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#seconds_until_auto_pause RdsDbCluster#seconds_until_auto_pause}
	SecondsUntilAutoPause *float64 `field:"optional" json:"secondsUntilAutoPause" yaml:"secondsUntilAutoPause"`
	// The action to take when the timeout is reached, either ForceApplyCapacityChange or RollbackCapacityChange.
	//
	// ForceApplyCapacityChange sets the capacity to the specified value as soon as possible.
	// RollbackCapacityChange, the default, ignores the capacity change if a scaling point isn't found in the timeout period.
	//
	// For more information, see Autoscaling for Aurora Serverless v1 in the Amazon Aurora User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#timeout_action RdsDbCluster#timeout_action}
	TimeoutAction *string `field:"optional" json:"timeoutAction" yaml:"timeoutAction"`
}

