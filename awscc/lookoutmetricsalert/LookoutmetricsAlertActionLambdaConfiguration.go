package lookoutmetricsalert


type LookoutmetricsAlertActionLambdaConfiguration struct {
	// ARN of a Lambda to send alert notifications to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_alert#lambda_arn LookoutmetricsAlert#lambda_arn}
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// ARN of an IAM role that LookoutMetrics should assume to access the Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_alert#role_arn LookoutmetricsAlert#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

