package s3objectlambdaaccesspoint


type S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3objectlambda_access_point#actions S3ObjectlambdaAccessPoint#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3objectlambda_access_point#content_transformation S3ObjectlambdaAccessPoint#content_transformation}.
	ContentTransformation *S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformation `field:"required" json:"contentTransformation" yaml:"contentTransformation"`
}

