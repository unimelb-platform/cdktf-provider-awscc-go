package sagemakerinferenceexperiment


type SagemakerInferenceExperimentDataStorageConfig struct {
	// The Amazon S3 bucket where the inference request and response data is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment#destination SagemakerInferenceExperiment#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Configuration specifying how to treat different headers.
	//
	// If no headers are specified SageMaker will by default base64 encode when capturing the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment#content_type SagemakerInferenceExperiment#content_type}
	ContentType *SagemakerInferenceExperimentDataStorageConfigContentType `field:"optional" json:"contentType" yaml:"contentType"`
	// The AWS Key Management Service key that Amazon SageMaker uses to encrypt captured data at rest using Amazon S3 server-side encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_inference_experiment#kms_key SagemakerInferenceExperiment#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

