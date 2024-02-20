package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupLifecycleHookSpecificationListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#lifecycle_hook_name AutoscalingAutoScalingGroup#lifecycle_hook_name}.
	LifecycleHookName *string `field:"required" json:"lifecycleHookName" yaml:"lifecycleHookName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#lifecycle_transition AutoscalingAutoScalingGroup#lifecycle_transition}.
	LifecycleTransition *string `field:"required" json:"lifecycleTransition" yaml:"lifecycleTransition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#default_result AutoscalingAutoScalingGroup#default_result}.
	DefaultResult *string `field:"optional" json:"defaultResult" yaml:"defaultResult"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#heartbeat_timeout AutoscalingAutoScalingGroup#heartbeat_timeout}.
	HeartbeatTimeout *float64 `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#notification_metadata AutoscalingAutoScalingGroup#notification_metadata}.
	NotificationMetadata *string `field:"optional" json:"notificationMetadata" yaml:"notificationMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#notification_target_arn AutoscalingAutoScalingGroup#notification_target_arn}.
	NotificationTargetArn *string `field:"optional" json:"notificationTargetArn" yaml:"notificationTargetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#role_arn AutoscalingAutoScalingGroup#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

