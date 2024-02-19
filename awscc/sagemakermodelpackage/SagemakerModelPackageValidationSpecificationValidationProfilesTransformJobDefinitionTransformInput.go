package sagemakermodelpackage


type SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput struct {
	// Describes the input source of a transform job and the way the transform job consumes it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#data_source SagemakerModelPackage#data_source}
	DataSource *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSource `field:"required" json:"dataSource" yaml:"dataSource"`
	// If your transform data is compressed, specify the compression type.
	//
	// Amazon SageMaker automatically decompresses the data for the transform job accordingly. The default value is None.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#compression_type SagemakerModelPackage#compression_type}
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The multipurpose internet mail extension (MIME) type of the data.
	//
	// Amazon SageMaker uses the MIME type with each http call to transfer data to the transform job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#content_type SagemakerModelPackage#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The method to use to split the transform job's data files into smaller batches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#split_type SagemakerModelPackage#split_type}
	SplitType *string `field:"optional" json:"splitType" yaml:"splitType"`
}

