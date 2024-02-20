package sagemakermodelpackage


type SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSourceS3DataSource struct {
	// The S3 Data Source Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#s3_data_type SagemakerModelPackage#s3_data_type}
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Depending on the value specified for the S3DataType, identifies either a key name prefix or a manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#s3_uri SagemakerModelPackage#s3_uri}
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

