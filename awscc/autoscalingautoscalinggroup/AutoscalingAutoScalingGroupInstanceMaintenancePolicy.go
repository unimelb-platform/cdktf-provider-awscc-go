package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupInstanceMaintenancePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#max_healthy_percentage AutoscalingAutoScalingGroup#max_healthy_percentage}.
	MaxHealthyPercentage *float64 `field:"optional" json:"maxHealthyPercentage" yaml:"maxHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#min_healthy_percentage AutoscalingAutoScalingGroup#min_healthy_percentage}.
	MinHealthyPercentage *float64 `field:"optional" json:"minHealthyPercentage" yaml:"minHealthyPercentage"`
}

