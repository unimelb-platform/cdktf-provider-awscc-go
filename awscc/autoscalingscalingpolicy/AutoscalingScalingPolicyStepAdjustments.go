package autoscalingscalingpolicy


type AutoscalingScalingPolicyStepAdjustments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#scaling_adjustment AutoscalingScalingPolicy#scaling_adjustment}.
	ScalingAdjustment *float64 `field:"required" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#metric_interval_lower_bound AutoscalingScalingPolicy#metric_interval_lower_bound}.
	MetricIntervalLowerBound *float64 `field:"optional" json:"metricIntervalLowerBound" yaml:"metricIntervalLowerBound"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#metric_interval_upper_bound AutoscalingScalingPolicy#metric_interval_upper_bound}.
	MetricIntervalUpperBound *float64 `field:"optional" json:"metricIntervalUpperBound" yaml:"metricIntervalUpperBound"`
}

