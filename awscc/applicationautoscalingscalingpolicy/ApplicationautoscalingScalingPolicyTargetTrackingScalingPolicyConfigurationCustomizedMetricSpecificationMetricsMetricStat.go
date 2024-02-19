package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStat struct {
	// The CloudWatch metric to return, including the metric name, namespace, and dimensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#metric ApplicationautoscalingScalingPolicy#metric}
	Metric *ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetric `field:"optional" json:"metric" yaml:"metric"`
	// The statistic to return. It can include any CloudWatch statistic or extended statistic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#stat ApplicationautoscalingScalingPolicy#stat}
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
	// The unit to use for the returned data points.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#unit ApplicationautoscalingScalingPolicy#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

