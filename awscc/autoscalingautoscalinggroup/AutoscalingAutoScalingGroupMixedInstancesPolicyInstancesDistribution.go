package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistribution struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#on_demand_allocation_strategy AutoscalingAutoScalingGroup#on_demand_allocation_strategy}.
	OnDemandAllocationStrategy *string `field:"optional" json:"onDemandAllocationStrategy" yaml:"onDemandAllocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#on_demand_base_capacity AutoscalingAutoScalingGroup#on_demand_base_capacity}.
	OnDemandBaseCapacity *float64 `field:"optional" json:"onDemandBaseCapacity" yaml:"onDemandBaseCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#on_demand_percentage_above_base_capacity AutoscalingAutoScalingGroup#on_demand_percentage_above_base_capacity}.
	OnDemandPercentageAboveBaseCapacity *float64 `field:"optional" json:"onDemandPercentageAboveBaseCapacity" yaml:"onDemandPercentageAboveBaseCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#spot_allocation_strategy AutoscalingAutoScalingGroup#spot_allocation_strategy}.
	SpotAllocationStrategy *string `field:"optional" json:"spotAllocationStrategy" yaml:"spotAllocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#spot_instance_pools AutoscalingAutoScalingGroup#spot_instance_pools}.
	SpotInstancePools *float64 `field:"optional" json:"spotInstancePools" yaml:"spotInstancePools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#spot_max_price AutoscalingAutoScalingGroup#spot_max_price}.
	SpotMaxPrice *string `field:"optional" json:"spotMaxPrice" yaml:"spotMaxPrice"`
}

