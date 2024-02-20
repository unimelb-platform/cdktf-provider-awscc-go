package lookoutmetricsalert


type LookoutmetricsAlertActionSnsConfiguration struct {
	// ARN of an IAM role that LookoutMetrics should assume to access the SNS topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_alert#role_arn LookoutmetricsAlert#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// ARN of an SNS topic to send alert notifications to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_alert#sns_topic_arn LookoutmetricsAlert#sns_topic_arn}
	SnsTopicArn *string `field:"required" json:"snsTopicArn" yaml:"snsTopicArn"`
}

