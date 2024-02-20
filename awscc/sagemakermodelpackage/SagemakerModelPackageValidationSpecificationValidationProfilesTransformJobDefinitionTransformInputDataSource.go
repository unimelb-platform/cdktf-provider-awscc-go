package sagemakermodelpackage


type SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSource struct {
	// Describes the S3 data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#s3_data_source SagemakerModelPackage#s3_data_source}
	S3DataSource *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSourceS3DataSource `field:"required" json:"s3DataSource" yaml:"s3DataSource"`
}

