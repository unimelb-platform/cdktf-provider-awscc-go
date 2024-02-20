package sagemakermodelqualityjobdefinition


type SagemakerModelQualityJobDefinitionJobResourcesClusterConfig struct {
	// The number of ML compute instances to use in the model monitoring job.
	//
	// For distributed processing jobs, specify a value greater than 1. The default value is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition#instance_count SagemakerModelQualityJobDefinition#instance_count}
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The ML compute instance type for the processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition#instance_type SagemakerModelQualityJobDefinition#instance_type}
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The size of the ML storage volume, in gigabytes, that you want to provision.
	//
	// You must specify sufficient ML storage for your scenario.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition#volume_size_in_gb SagemakerModelQualityJobDefinition#volume_size_in_gb}
	VolumeSizeInGb *float64 `field:"required" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition#volume_kms_key_id SagemakerModelQualityJobDefinition#volume_kms_key_id}
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
}

