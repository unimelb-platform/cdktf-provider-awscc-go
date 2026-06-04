package autoscalingscalingpolicy


type AutoscalingScalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/autoscaling_scaling_policy#dimensions AutoscalingScalingPolicy#dimensions}.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/autoscaling_scaling_policy#metric_name AutoscalingScalingPolicy#metric_name}.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/autoscaling_scaling_policy#namespace AutoscalingScalingPolicy#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

