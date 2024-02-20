package cassandratable


type CassandraTableAutoScalingSpecificationsReadCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#target_value CassandraTable#target_value}.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#disable_scale_in CassandraTable#disable_scale_in}.
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#scale_in_cooldown CassandraTable#scale_in_cooldown}.
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#scale_out_cooldown CassandraTable#scale_out_cooldown}.
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

