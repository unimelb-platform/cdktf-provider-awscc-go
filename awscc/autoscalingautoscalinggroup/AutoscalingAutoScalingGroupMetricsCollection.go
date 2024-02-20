package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMetricsCollection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#granularity AutoscalingAutoScalingGroup#granularity}.
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#metrics AutoscalingAutoScalingGroup#metrics}.
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
}

