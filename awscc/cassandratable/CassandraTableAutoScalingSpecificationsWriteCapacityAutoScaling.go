package cassandratable


type CassandraTableAutoScalingSpecificationsWriteCapacityAutoScaling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#auto_scaling_disabled CassandraTable#auto_scaling_disabled}.
	AutoScalingDisabled interface{} `field:"optional" json:"autoScalingDisabled" yaml:"autoScalingDisabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#maximum_units CassandraTable#maximum_units}.
	MaximumUnits *float64 `field:"optional" json:"maximumUnits" yaml:"maximumUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#minimum_units CassandraTable#minimum_units}.
	MinimumUnits *float64 `field:"optional" json:"minimumUnits" yaml:"minimumUnits"`
	// Represents scaling policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#scaling_policy CassandraTable#scaling_policy}
	ScalingPolicy *CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicy `field:"optional" json:"scalingPolicy" yaml:"scalingPolicy"`
}

