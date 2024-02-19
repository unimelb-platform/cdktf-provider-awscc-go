package s3objectlambdaaccesspoint


type S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambda struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3objectlambda_access_point#function_arn S3ObjectlambdaAccessPoint#function_arn}.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3objectlambda_access_point#function_payload S3ObjectlambdaAccessPoint#function_payload}.
	FunctionPayload *string `field:"optional" json:"functionPayload" yaml:"functionPayload"`
}

