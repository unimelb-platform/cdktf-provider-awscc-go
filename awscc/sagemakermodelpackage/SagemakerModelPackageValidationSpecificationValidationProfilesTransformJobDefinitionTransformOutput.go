package sagemakermodelpackage


type SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput struct {
	// The Amazon S3 path where you want Amazon SageMaker to store the results of the transform job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#s3_output_path SagemakerModelPackage#s3_output_path}
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// The MIME type used to specify the output data.
	//
	// Amazon SageMaker uses the MIME type with each http call to transfer data from the transform job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#accept SagemakerModelPackage#accept}
	Accept *string `field:"optional" json:"accept" yaml:"accept"`
	// Defines how to assemble the results of the transform job as a single S3 object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#assemble_with SagemakerModelPackage#assemble_with}
	AssembleWith *string `field:"optional" json:"assembleWith" yaml:"assembleWith"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#kms_key_id SagemakerModelPackage#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

