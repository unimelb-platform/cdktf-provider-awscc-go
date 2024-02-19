package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyStepScalingPolicyConfigurationStepAdjustments struct {
	// The amount by which to scale, based on the specified adjustment type.
	//
	// A positive value adds to the current capacity while a negative number removes from the current capacity. For exact capacity, you must specify a positive value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#scaling_adjustment ApplicationautoscalingScalingPolicy#scaling_adjustment}
	ScalingAdjustment *float64 `field:"required" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// The lower bound for the difference between the alarm threshold and the CloudWatch metric.
	//
	// If the metric value is above the breach threshold, the lower bound is inclusive (the metric must be greater than or equal to the threshold plus the lower bound). Otherwise, it is exclusive (the metric must be greater than the threshold plus the lower bound). A null value indicates negative infinity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#metric_interval_lower_bound ApplicationautoscalingScalingPolicy#metric_interval_lower_bound}
	MetricIntervalLowerBound *float64 `field:"optional" json:"metricIntervalLowerBound" yaml:"metricIntervalLowerBound"`
	// The upper bound for the difference between the alarm threshold and the CloudWatch metric.
	//
	// If the metric value is above the breach threshold, the upper bound is exclusive (the metric must be less than the threshold plus the upper bound). Otherwise, it is inclusive (the metric must be less than or equal to the threshold plus the upper bound). A null value indicates positive infinity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#metric_interval_upper_bound ApplicationautoscalingScalingPolicy#metric_interval_upper_bound}
	MetricIntervalUpperBound *float64 `field:"optional" json:"metricIntervalUpperBound" yaml:"metricIntervalUpperBound"`
}

