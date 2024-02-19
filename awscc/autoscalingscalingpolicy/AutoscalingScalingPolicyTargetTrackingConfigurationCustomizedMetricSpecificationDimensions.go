package autoscalingscalingpolicy


type AutoscalingScalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationDimensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#name AutoscalingScalingPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_scaling_policy#value AutoscalingScalingPolicy#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

