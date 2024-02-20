package s3bucket


type S3BucketNotificationConfigurationLambdaConfigurations struct {
	// The Amazon S3 bucket event for which to invoke the AWS Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#event S3Bucket#event}
	Event *string `field:"required" json:"event" yaml:"event"`
	// The Amazon Resource Name (ARN) of the AWS Lambda function that Amazon S3 invokes when the specified event type occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#function S3Bucket#function}
	Function *string `field:"required" json:"function" yaml:"function"`
	// The filtering rules that determine which objects invoke the AWS Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#filter S3Bucket#filter}
	Filter *S3BucketNotificationConfigurationLambdaConfigurationsFilter `field:"optional" json:"filter" yaml:"filter"`
}

