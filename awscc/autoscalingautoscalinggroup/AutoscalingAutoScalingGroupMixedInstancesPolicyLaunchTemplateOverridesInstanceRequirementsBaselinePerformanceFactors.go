package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsBaselinePerformanceFactors struct {
	// The CPU performance to consider, using an instance family as the baseline reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/autoscaling_auto_scaling_group#cpu AutoscalingAutoScalingGroup#cpu}
	Cpu *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsBaselinePerformanceFactorsCpu `field:"optional" json:"cpu" yaml:"cpu"`
}

