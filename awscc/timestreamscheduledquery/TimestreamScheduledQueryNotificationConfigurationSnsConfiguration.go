package timestreamscheduledquery


type TimestreamScheduledQueryNotificationConfigurationSnsConfiguration struct {
	// SNS topic ARN that the scheduled query status notifications will be sent to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/timestream_scheduled_query#topic_arn TimestreamScheduledQuery#topic_arn}
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

