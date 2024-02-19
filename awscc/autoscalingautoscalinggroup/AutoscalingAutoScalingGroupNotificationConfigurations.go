package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupNotificationConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#topic_arn AutoscalingAutoScalingGroup#topic_arn}.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#notification_types AutoscalingAutoScalingGroup#notification_types}.
	NotificationTypes *[]*string `field:"optional" json:"notificationTypes" yaml:"notificationTypes"`
}

