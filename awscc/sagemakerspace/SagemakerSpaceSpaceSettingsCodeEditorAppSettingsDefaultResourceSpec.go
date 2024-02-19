package sagemakerspace


type SagemakerSpaceSpaceSettingsCodeEditorAppSettingsDefaultResourceSpec struct {
	// The instance type that the image version runs on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#instance_type SagemakerSpace#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ARN of the SageMaker image that the image version belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#sage_maker_image_arn SagemakerSpace#sage_maker_image_arn}
	SageMakerImageArn *string `field:"optional" json:"sageMakerImageArn" yaml:"sageMakerImageArn"`
	// The ARN of the image version created on the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#sage_maker_image_version_arn SagemakerSpace#sage_maker_image_version_arn}
	SageMakerImageVersionArn *string `field:"optional" json:"sageMakerImageVersionArn" yaml:"sageMakerImageVersionArn"`
}

