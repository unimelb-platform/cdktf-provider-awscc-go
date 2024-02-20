package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification struct {
	// The metric type. The ALBRequestCountPerTarget metric type applies only to Spot Fleets and ECS services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#predefined_metric_type ApplicationautoscalingScalingPolicy#predefined_metric_type}
	PredefinedMetricType *string `field:"required" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// Identifies the resource associated with the metric type.
	//
	// You can't specify a resource label unless the metric type is ALBRequestCountPerTarget and there is a target group attached to the Spot Fleet or ECS service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationautoscaling_scaling_policy#resource_label ApplicationautoscalingScalingPolicy#resource_label}
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

