package lambdafunction


type LambdaFunctionDeadLetterConfig struct {
	// The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#target_arn LambdaFunction#target_arn}
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

