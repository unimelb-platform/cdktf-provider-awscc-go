package autoscalingscalingpolicy


type AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsCustomizedScalingMetricSpecificationMetricDataQueriesMetricStat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#metric AutoscalingScalingPolicy#metric}.
	Metric *AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsCustomizedScalingMetricSpecificationMetricDataQueriesMetricStatMetric `field:"required" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#stat AutoscalingScalingPolicy#stat}.
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#unit AutoscalingScalingPolicy#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

