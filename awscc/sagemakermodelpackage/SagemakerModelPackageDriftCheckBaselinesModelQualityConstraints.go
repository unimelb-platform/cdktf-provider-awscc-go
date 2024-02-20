package sagemakermodelpackage


type SagemakerModelPackageDriftCheckBaselinesModelQualityConstraints struct {
	// The type of content stored in the metric source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#content_type SagemakerModelPackage#content_type}
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// The Amazon S3 URI for the metric source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#s3_uri SagemakerModelPackage#s3_uri}
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The digest of the metric source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#content_digest SagemakerModelPackage#content_digest}
	ContentDigest *string `field:"optional" json:"contentDigest" yaml:"contentDigest"`
}

