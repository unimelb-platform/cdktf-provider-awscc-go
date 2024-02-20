package sagemakeruserprofile


type SagemakerUserProfileTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#key SagemakerUserProfile#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#value SagemakerUserProfile#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

