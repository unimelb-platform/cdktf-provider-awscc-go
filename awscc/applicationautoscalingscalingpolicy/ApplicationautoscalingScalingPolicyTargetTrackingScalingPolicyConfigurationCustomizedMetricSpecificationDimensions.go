package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions struct {
	// The name of the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#name ApplicationautoscalingScalingPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#value ApplicationautoscalingScalingPolicy#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

