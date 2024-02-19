package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyStepScalingPolicyConfiguration struct {
	// Specifies how the ScalingAdjustment value in a StepAdjustment is interpreted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#adjustment_type ApplicationautoscalingScalingPolicy#adjustment_type}
	AdjustmentType *string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// The amount of time, in seconds, to wait for a previous scaling activity to take effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#cooldown ApplicationautoscalingScalingPolicy#cooldown}
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// The aggregation type for the CloudWatch metrics.
	//
	// Valid values are Minimum, Maximum, and Average. If the aggregation type is null, the value is treated as Average
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#metric_aggregation_type ApplicationautoscalingScalingPolicy#metric_aggregation_type}
	MetricAggregationType *string `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// The minimum value to scale by when the adjustment type is PercentChangeInCapacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#min_adjustment_magnitude ApplicationautoscalingScalingPolicy#min_adjustment_magnitude}
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// A set of adjustments that enable you to scale based on the size of the alarm breach.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#step_adjustments ApplicationautoscalingScalingPolicy#step_adjustments}
	StepAdjustments interface{} `field:"optional" json:"stepAdjustments" yaml:"stepAdjustments"`
}

