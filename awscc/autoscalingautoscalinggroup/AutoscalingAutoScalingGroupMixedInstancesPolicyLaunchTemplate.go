package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplate struct {
	// The launch template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/autoscaling_auto_scaling_group#launch_template_specification AutoscalingAutoScalingGroup#launch_template_specification}
	LaunchTemplateSpecification *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Any properties that you specify override the same properties in the launch template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/autoscaling_auto_scaling_group#overrides AutoscalingAutoScalingGroup#overrides}
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

