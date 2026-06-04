package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupAvailabilityZoneDistribution struct {
	// If launches fail in an Availability Zone, the following strategies are available.
	//
	// The default is ``balanced-best-effort``.
	//   +   ``balanced-only`` - If launches fail in an Availability Zone, Auto Scaling will continue to attempt to launch in the unhealthy zone to preserve a balanced distribution.
	//   +   ``balanced-best-effort`` - If launches fail in an Availability Zone, Auto Scaling will attempt to launch in another healthy Availability Zone instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/autoscaling_auto_scaling_group#capacity_distribution_strategy AutoscalingAutoScalingGroup#capacity_distribution_strategy}
	CapacityDistributionStrategy *string `field:"optional" json:"capacityDistributionStrategy" yaml:"capacityDistributionStrategy"`
}

