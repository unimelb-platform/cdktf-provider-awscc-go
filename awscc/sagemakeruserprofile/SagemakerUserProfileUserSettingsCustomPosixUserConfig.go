package sagemakeruserprofile


type SagemakerUserProfileUserSettingsCustomPosixUserConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#gid SagemakerUserProfile#gid}.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#uid SagemakerUserProfile#uid}.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
}

