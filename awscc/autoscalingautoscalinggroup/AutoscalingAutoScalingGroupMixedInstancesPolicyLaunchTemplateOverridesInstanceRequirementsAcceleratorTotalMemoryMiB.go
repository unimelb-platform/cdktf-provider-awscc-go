package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsAcceleratorTotalMemoryMiB struct {
	// The memory maximum in MiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/autoscaling_auto_scaling_group#max AutoscalingAutoScalingGroup#max}
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The memory minimum in MiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/autoscaling_auto_scaling_group#min AutoscalingAutoScalingGroup#min}
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

