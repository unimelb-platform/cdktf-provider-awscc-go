package lambdafunction


type LambdaFunctionCode struct {
	// ImageUri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#image_uri LambdaFunction#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// An Amazon S3 bucket in the same AWS Region as your function.
	//
	// The bucket can be in a different AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#s3_bucket LambdaFunction#s3_bucket}
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The Amazon S3 key of the deployment package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#s3_key LambdaFunction#s3_key}
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
	// For versioned objects, the version of the deployment package object to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#s3_object_version LambdaFunction#s3_object_version}
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// The source code of your Lambda function.
	//
	// If you include your function source inline with this parameter, AWS CloudFormation places it in a file named index and zips it to create a deployment package..
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#zip_file LambdaFunction#zip_file}
	ZipFile *string `field:"optional" json:"zipFile" yaml:"zipFile"`
}

