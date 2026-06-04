package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsCustomizedCapacityMetricSpecification struct {
	// One or more metric data queries to provide data points for a metric specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationautoscaling_scaling_policy#metric_data_queries ApplicationautoscalingScalingPolicy#metric_data_queries}
	MetricDataQueries interface{} `field:"optional" json:"metricDataQueries" yaml:"metricDataQueries"`
}

