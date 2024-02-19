package autoscalingscalingpolicy


type AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsCustomizedScalingMetricSpecificationMetricDataQueries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#id AutoscalingScalingPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#expression AutoscalingScalingPolicy#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#label AutoscalingScalingPolicy#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#metric_stat AutoscalingScalingPolicy#metric_stat}.
	MetricStat *AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsCustomizedScalingMetricSpecificationMetricDataQueriesMetricStat `field:"optional" json:"metricStat" yaml:"metricStat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#return_data AutoscalingScalingPolicy#return_data}.
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

