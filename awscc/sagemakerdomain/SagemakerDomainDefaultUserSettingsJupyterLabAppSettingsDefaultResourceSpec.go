package sagemakerdomain


type SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpec struct {
	// The instance type that the image version runs on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#instance_type SagemakerDomain#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The Amazon Resource Name (ARN) of the Lifecycle Configuration to attach to the Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#lifecycle_config_arn SagemakerDomain#lifecycle_config_arn}
	LifecycleConfigArn *string `field:"optional" json:"lifecycleConfigArn" yaml:"lifecycleConfigArn"`
	// The Amazon Resource Name (ARN) of the SageMaker image that the image version belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#sage_maker_image_arn SagemakerDomain#sage_maker_image_arn}
	SageMakerImageArn *string `field:"optional" json:"sageMakerImageArn" yaml:"sageMakerImageArn"`
	// The Amazon Resource Name (ARN) of the image version created on the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#sage_maker_image_version_arn SagemakerDomain#sage_maker_image_version_arn}
	SageMakerImageVersionArn *string `field:"optional" json:"sageMakerImageVersionArn" yaml:"sageMakerImageVersionArn"`
}

