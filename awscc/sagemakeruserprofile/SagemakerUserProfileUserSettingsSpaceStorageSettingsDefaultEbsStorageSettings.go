package sagemakeruserprofile


type SagemakerUserProfileUserSettingsSpaceStorageSettingsDefaultEbsStorageSettings struct {
	// Default size of the Amazon EBS volume in Gb.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#default_ebs_volume_size_in_gb SagemakerUserProfile#default_ebs_volume_size_in_gb}
	DefaultEbsVolumeSizeInGb *float64 `field:"required" json:"defaultEbsVolumeSizeInGb" yaml:"defaultEbsVolumeSizeInGb"`
	// Maximum size of the Amazon EBS volume in Gb. Must be greater than or equal to the DefaultEbsVolumeSizeInGb.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#maximum_ebs_volume_size_in_gb SagemakerUserProfile#maximum_ebs_volume_size_in_gb}
	MaximumEbsVolumeSizeInGb *float64 `field:"required" json:"maximumEbsVolumeSizeInGb" yaml:"maximumEbsVolumeSizeInGb"`
}

