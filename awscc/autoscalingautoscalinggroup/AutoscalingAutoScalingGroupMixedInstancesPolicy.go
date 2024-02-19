package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#launch_template AutoscalingAutoScalingGroup#launch_template}.
	LaunchTemplate *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplate `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#instances_distribution AutoscalingAutoScalingGroup#instances_distribution}.
	InstancesDistribution *AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistribution `field:"optional" json:"instancesDistribution" yaml:"instancesDistribution"`
}

