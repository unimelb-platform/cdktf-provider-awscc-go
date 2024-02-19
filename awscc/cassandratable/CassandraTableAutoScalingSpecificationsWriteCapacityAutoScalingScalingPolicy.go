package cassandratable


type CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicy struct {
	// Represents configuration for target tracking scaling policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#target_tracking_scaling_policy_configuration CassandraTable#target_tracking_scaling_policy_configuration}
	TargetTrackingScalingPolicyConfiguration *CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyTargetTrackingScalingPolicyConfiguration `field:"optional" json:"targetTrackingScalingPolicyConfiguration" yaml:"targetTrackingScalingPolicyConfiguration"`
}

