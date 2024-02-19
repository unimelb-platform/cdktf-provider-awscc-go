package autoscalingscalingpolicy


type AutoscalingScalingPolicyTargetTrackingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#target_value AutoscalingScalingPolicy#target_value}.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#customized_metric_specification AutoscalingScalingPolicy#customized_metric_specification}.
	CustomizedMetricSpecification *AutoscalingScalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification `field:"optional" json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#disable_scale_in AutoscalingScalingPolicy#disable_scale_in}.
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#predefined_metric_specification AutoscalingScalingPolicy#predefined_metric_specification}.
	PredefinedMetricSpecification *AutoscalingScalingPolicyTargetTrackingConfigurationPredefinedMetricSpecification `field:"optional" json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
}

