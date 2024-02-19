package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification struct {
	// The dimensions of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#dimensions ApplicationautoscalingScalingPolicy#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The name of the metric.
	//
	// To get the exact metric name, namespace, and dimensions, inspect the Metric object that is returned by a call to ListMetrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#metric_name ApplicationautoscalingScalingPolicy#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The metrics to include in the target tracking scaling policy, as a metric data query.
	//
	// This can include both raw metric and metric math expressions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#metrics ApplicationautoscalingScalingPolicy#metrics}
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// The namespace of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#namespace ApplicationautoscalingScalingPolicy#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The statistic of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#statistic ApplicationautoscalingScalingPolicy#statistic}
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// The unit of the metric.
	//
	// For a complete list of the units that CloudWatch supports, see the MetricDatum data type in the Amazon CloudWatch API Reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#unit ApplicationautoscalingScalingPolicy#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

